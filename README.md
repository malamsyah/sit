# Service Isolation Test
[![build](https://github.com/malamsyah/sit/actions/workflows/main.yml/badge.svg)](https://github.com/malamsyah/sit/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/malamsyah/sit)](https://goreportcard.com/report/github.com/malamsyah/sit)

A service-isolation-test library built in Golang

## Dependencies
* Wiremock running at port 8080
    ```
    docker run -it --rm -p 8080:8080 wiremock/wiremock
    ```
* Install `gotestsum`
    ```
    go install gotest.tools/gotestsum@latest
    ```
* Running unit tests
    ```
    make test
    ```

## Documentation
🚧 WIP

## Example
🚧 WIP