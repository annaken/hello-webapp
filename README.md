# hello-webapp

This is a demo webapp which implements some simple functionality and exposes a metrics page for consumption by Prometheus.

Files included in this repo are intended to build a container using a Dockerfile, and deploy to kubernetes.

## Example deployment flow via CICD:

* Source code is updated in the git repo and a PR is submitted
* The PR initiates a CICD pipeline with the following steps:
  * run linting and static code analysis
  * build dev image
  * deploy dev image to the development environment
  * run tests (unit, integration, regression, smoke) (discuss: parallel or serial)
  * if tests pass, approve PR
  * if the tests fail, notify the developer
* When the PR is approved, the app can either be automatically released to the production environment, or marked as prod-ready and await manual intervention
* When the app is released to prod, monitoring should alert quickly if there is a problem
* If there is an issue, it might be appropriate to implement an automatic rollback


