# app-cluster-api
​
The application cluster api is the component responsible of receiving request from the management cluster
components, and redirecting those requests to the appropriate internal components.

The internal structure follows a gRPC service structure where the name of the Service matches the receiver
name to facilitate implementation and tracking.
​
## Getting Started
​
This component must be deployed and running on an application cluster for it to be usable by the management cluster.
​
### Prerequisites
​
* Downstream components on the application cluster must be running.
* The provisioning process executed by the `provisioner` on the managemenet cluster installs [cert-manager](https://github.com/jetstack/cert-manager),
and it creates the client certificate to be installed on the associated ingress.
* The installation process must create an ingress using the aforementioned certificate to validate that the management cluster certificate is issued
by the configured CA.
​
### Build and compile
​
In order to build and compile this repository use the provided Makefile:
​
```
make all
```
​
This operation generates the binaries for this repo, download dependencies,
run existing tests and generate ready-to-deploy Kubernetes files.
​
### Run tests
​
Tests are executed using Ginkgo. To run all the available tests:
​
```
make test
```
​
### Update dependencies
​
Dependencies are managed using Godep. For an automatic dependencies download use:
​
```
make dep
```
​
In order to have all dependencies up-to-date run:
​
```
dep ensure -update -v
```
​
## Known Issues
​
* The authentication of the management cluster if fully delegated to the ingress, future versions will include tighter
control on which management clusters can send commands to the application cluster.
​
​
## Contributing
​
Please read [contributing.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.
​
## Versioning
​
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/app-cluster-api/tags). 
​
## Authors
​
See also the list of [contributors](https://github.com/nalej/app-cluster-api/contributors) who participated in this project.
​
## License
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.