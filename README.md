# hello-webapp

This is a demo webapp which implements some simple functionality and exposes a metrics page for consumption by Prometheus.

Files included in this repo are intended to build a container using a Dockerfile, and deploy to kubernetes.

## Deploy to development environment:

* Source code is updated in the git repo and a PR is submitted
* Linting and static code analysis is run on the code
* Image is built
* Image is tested (unit, regression)
* Image is tagged as dev-release
* Dev image is deployed to the development environment
* Deployment is tested (run integration, smoke, load tests)
* If any stage of testing fails, block the PR
* If all tests pass, approve the PR

## Deploy to production environment:

* PR is merged
* EITHER: image is tagged as prod-release and awaits manual release
* OR: image is automatically released to the production account
* Monitoring/alerting should notify if there is an issue
* Rollback plan should be ready
* Consider automatic rollback if certain critieria are met

