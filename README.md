# Terraform IP Address Manager (tf-ipam)

This repository contains the source code for the tf-ipam provider, a Terraform provider for IP Address Management.

- `examples/` contains helpful examples to get you started
- `internal/` contains the source source for the provider

This provider requires a backend for IPAM information to be stored. While I sought out to store everything in Terraform's state, I found limitations within Terraform that prevented this from being a reality with parent-child resources like this provider implements. As a result, having a backend to store this information was a necessity. I have made an attempt to have a few commonly used backends built in. If there's other backends that would be useful, please open a GitHub issue with your suggestions. I also welcome PR's.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install .
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

TODO: Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install .`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

```shell
make testacc
```
