# go.formulabun.club

[formulabun.club]() is an srb2kart server. It turned into a coding/histing project which includes go.formulabun.club.

This [go.formulabun.club]() repository contains all the packages written for formulabun. From utility functions in [functional](https://godoc.formulabun.club/pkg/go.formulabun.club/functional/), to all [srb2kart](https://godoc.formulabun.club/pkg/go.formulabun.club/srb2kart/) related code and services like the [srb2kart socket to json translator](https://godoc.formulabun.club/pkg/go.formulabun.club/translator/).

## Technologies used

Everything is written in [Go](https://go.dev/), using the teachings form [gopl](http://www.gopl.io/). It is a simple language but draws a lot of power from that simplicity. [formulabun.club]() runs on a unix [VPS](https://en.wikipedia.org/wiki/Virtual_private_server) from [contabo](https://contabo.com/en/), a very good price/performance host. It runs [nginx](https://nginx.org/en/) with a cercificate from [Let's Encrypt](https://letsencrypt.org/).

The deployment on the VPS is done using [docker](https://www.docker.com/) and [docker compose](https://docs.docker.com/compose/). This allows for a [microsevice](https://microservices.io/) like deployment which keeps each service simple, going hand in hand with the simplicity of the Go language. In order to avoid writing the boilerplate HTTP server and client code, the API's are defined in [openapi](https://www.openapis.org/), from which the server and client code is generated using the Go [openapi generators](https://openapi-generator.tech/). A process that happens only a few times per service, due to their simplicity and small size. Both the docker compose files and the api specifications are written in [YAML](https://yaml.org/).

## Setting up a development environment

First of, **DO NOT CLONE THIS REPOSITORY.** This is only a collection of the packages and only serves the godoc site on [godoc.formulabun.club]() and contains build and release docker compose files. It is not meant for development.

Instead, clone the packages (the submodules in this repository) you want to work on in `GOPATH/go/src/go.formulabun.club/`. 

If this is the only package you're working in you're already good to go. Pick an issue, make a branch and create a pull request when you're done.

If you're working in multiple package at once, clone them all and use a [replace directive](https://go.dev/doc/modules/gomod-ref#replace) in the necessary `go.mod` files. These will generally take the form of `go.formulabun.club/<package> => ../<package>`.

Because Go imports can't form cycles, you can complete the pull requests for each package in order. Do this without breaking other packages. Avoid writing breaking code or having to version a package. This may be a challenge but you should write code that's isn't prone to breaking changes anyway 😉. Think of the future and write wisely.

### Running the development environment

See the docker files ig. Still working on that tbh.