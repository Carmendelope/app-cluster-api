/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package server

import (
	"context"
	"fmt"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SSLClientSubjectDN with the key of the header send by the ingress controller with the DN of the client
const SSLClientSubjectDN = "ssl-client-subject-dn"

// SSLClientVerify with the key of the header send by the ingress controller with the verification result
const SSLClientVerify = "ssl-client-verify"

// SSLClientVerifyOk with a constant used by the ingress controller in case the certificate is valid.
const SSLClientVerifyOk = "SUCCESS"

// WithClientCertValidator creates a server option that encapsulates an interceptor that validates
// that the incoming client matches the configured management public host.
func WithClientCertValidator(config *Config) grpc.ServerOption {
	return grpc.UnaryInterceptor(ClientCertValidator(config))
}

// ClientCertValidator creates a server interceptor that checks if the client certificate matches
// the configured management public host.
func ClientCertValidator(config *Config) grpc.UnaryServerInterceptor {
	// build the CN entry as found on the certificate.
	managementCN := fmt.Sprintf("CN=*.%s", config.ManagementPublicHost)

	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		// Check metadata is present.
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Error().Msg("no metadata found on incoming request")
			return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("unauthorized request").WithParams(info.FullMethod))
		}
		// Check the certificate has already been verified by the ingress controller.
		verified, found := md[SSLClientVerify]
		if !found {
			log.Error().Msg("SSL verified not found")
			return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("unauthorized request").WithParams(info.FullMethod))
		}
		if len(verified) != 1 || verified[0] != SSLClientVerifyOk {
			log.Error().Strs("verified", verified).Msg("SSL not verified")
			return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("unauthorized request").WithParams(info.FullMethod))
		}
		// Check the subject of the incoming request
		clientDN, found := md[SSLClientSubjectDN]
		if !found {
			log.Error().Msg("SSL Client DN not found")
			return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("unauthorized request").WithParams(info.FullMethod))
		}
		if len(clientDN) != 1 || clientDN[0] != managementCN {
			log.Error().Strs("verified", clientDN).Msg("Invalid management CN")
			return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("unauthorized request").WithParams(info.FullMethod))
		}
		// All correct, dispatch the request.
		return handler(ctx, req)
	}
}
