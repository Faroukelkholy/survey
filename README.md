# Survey


## Table of contents

1. [Documentation](#Documentation)
2. [Run](#Run)
3. [Design](#Design)
4. [Testing](#Testing)
5. [Linting](#Linting)
6. [Production_Ready](#Production_Ready)

## 1. Documentation

> IDL for the http api in swagger resides at ./docs folder.

## 2. Run

Project is created with:

* Go
* Echo
* Mongo
* Docker
* Mockery

```shell
make docker-run
```

> Note: The service will be running on port 3000 and the DB will be seeded 

## 3. Design

* I rely on SOLID principles to create code that loosely coupled and easy to extend.

## 4. Testing

```shell
make test
```

## 5. Linting

Project uses [golangci-lint](https://golangci-lint.run/). It is a go linter aggregator that can enable up to 48 linters.

#### 5.1 Configuartion

golanci-lint configuration is found in .golangci.yml file.

#### 5.2 Installation

```
# binary installation for linux and Windows, binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2
```

Check if the tool is corectly installed

```
golangci-lint --version
```

#### 5.3 Run the tool with the enabled linter

```shell
make golangci
```

golangci-lint print out warning messages in command-line related to the enabled linters in order to fix them.

#### 5.4 Linters commands to automatically fix warning messages provided

To format all files in the project based on the gofmt linter. [Ref](https://stackoverflow.com/a/13333931/5486622)

```shell
make gofmt
```

To fix go import packages linting warnings based on goimport linter. [Ref](https://stackoverflow.com/a/59964885/5486622)

```shell
make goimport
```

[Guide](https://stackoverflow.com/a/38714480/5486622) How you should group your package based on golang structure.

to run both linters

```shell
make lint
```

## 6. Production_Ready

### Configurable
Any service must be configurable:

I used sample module godotenv for this poc but in production, I would use something like Viper, it can read more complex configuration from yaml file, command-line flag parsing and also env variable.

### Security

* Service should be running over HTTPS or at least the api gateway or reverse proxy should handle requests outside LAN over HTTPS.
* Service authorize requests using jwt token.

### Testing
* Implement more coverage with Unit Test, targeting all edge cases.
* Implement integration tests and e2e tests

### Pre-commit
* Create a pre-commit hook to run linters and tests to make sure the code is free from lint warnings and test errors.

### CI
* Create a CI pipeline will run linters, run unit tests, build code.

### Deployment
* The app should be deployed using Kubernetes that will handle scaling the service as needed, service registry, reverse
  proxy and load balancer.

### Logging, Monitoring and Tracing

* Logging better in my service, provide info and context for stages in code. Importantly, logging warnings and errors to detect anomalies and close bugs qualitatively.
* Use ELK to be able to search inside logs, k8s metrics
* Use Jaeger to create spans and metrics for my service




