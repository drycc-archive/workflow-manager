
# Drycc Workflow Manager

[![Build Status](https://travis-ci.org/drycc/workflow-manager.svg?branch=master)](https://travis-ci.org/drycc/workflow-manager)
[![codecov](https://codecov.io/gh/drycc/workflow-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/drycc/workflow-manager)
[![Go Report Card](https://goreportcard.com/badge/github.com/drycc/workflow-manager)](https://goreportcard.com/report/github.com/drycc/workflow-manager)
[![codebeat badge](https://codebeat.co/badges/9bb37e2b-b7e5-4c70-9cbc-ddd26632e394)](https://codebeat.co/projects/github-com-drycc-workflow-manager-master)
[![Docker Repository on Quay](https://quay.io/repository/drycc/workflow-manager/status "Docker Repository on Quay")](https://quay.io/repository/drycc/workflow-manager)

This repository contains the manager component for Drycc Workflow. Drycc
(pronounced DAY-iss) Workflow is an open source Platform as a Service (PaaS)
that adds a developer-friendly layer to any [Kubernetes][k8s-home] cluster,
making it easy to deploy and manage applications on your own servers.

For more information about Drycc Workflow, please visit the main project page at
https://github.com/drycc/workflow.

We welcome your input! If you have feedback on Workflow Manager,
please [submit an issue][issues]. If you'd like to participate in development,
please read the "Development" section below and [submit a pull request][prs].

## Stay up to date

One of the primary goals for Workflow Manager is notifying operators of
component freshness. Workflow Manager will regularly check your cluster against
the latest stable components. If components are missing due to failure or are
simply out of date, Workflow operators will know at a glance.

By default, Workflow Manager will make version checks to an external service.
This submits component and version information to our versions service running
at [https://versions.drycc.com](https://versions.drycc.com). If you prefer this
check not happen, you may disable the function by setting
`WORKFLOW_MANAGER_CHECKVERSIONS` to `false` in the Workflow Manager's
Replication Controller.

## Workflow Doctor

Deployed closest to any potential problem, Workflow Manager is also designed to
help when things aren't going well. To aid troubleshooting efforts cluster
operators will be able to easily gather and securely submit cluster health and
status information to the Drycc team.

Functionality will be added in a later release.

# Development

The Drycc project welcomes contributions from all developers. The high level
process for development matches many other open source projects. See below for
an outline.

* Fork this repository
* Make your changes
* [Submit a pull request][prs] (PR) to this repository with your changes, and unit tests whenever possible
    * If your PR fixes any [issues][issues], make sure you write `Fixes #1234` in your PR description (where `#1234` is the number of the issue you're closing)
* The Drycc core contributors will review your code. After each of them sign off on your code, they'll label your PR with `LGTM1` and `LGTM2` (respectively). Once that happens, a contributor will merge it

## Docker Based Development Environment

The preferred environment for development uses [the `go-dev` Docker
image](https://github.com/drycc/docker-go-dev). The tools described in this
section are used to build, test, package and release each version of Drycc.

To use it yourself, you must have [make](https://www.gnu.org/software/make/)
installed and Docker installed and running on your local development machine.

If you don't have Docker installed, please go to https://www.docker.com/ to
install it.

After you have those dependencies, bootstrap dependencies with `make bootstrap`,
build your code with `make build` and execute unit tests with `make test`.

## Native Go Development Environment

You can also use the standard `go` toolchain to build and test if you prefer.
To do so, you'll need [glide](https://github.com/Masterminds/glide) 0.9 or
above and [Go 1.6](http://golang.org) or above installed.

After you have those dependencies, you can build and unit-test your code with
`go build` and `go test $(glide nv)`, respectively.

Note that you will not be able to build or push Docker images using this method
of development.

# Testing

The Drycc project requires that as much code as possible is unit tested, but the
core contributors also recognize that some code must be tested at a higher
level (functional or integration tests, for example).


[issues]: https://github.com/drycc/workflow-manager/issues
[prs]: https://github.com/drycc/workflow-manager/pulls
[k8s-home]: https://kubernetes.io
[v2.18]: https://github.com/drycc/workflow/releases/tag/v2.18.0
