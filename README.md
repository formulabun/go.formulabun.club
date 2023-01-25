# go.formulabun.club

[formulabun.club]() is an srb2kart server. It turned into a coding/histing project which includes go.formulabun.club.

This [go.formulabun.club]() repository contains all the packages written for formulabun. From utility functions in [functional](https://godoc.formulabun.club/pkg/go.formulabun.club/functional/), to all [srb2kart](https://godoc.formulabun.club/pkg/go.formulabun.club/srb2kart/) related code and services like the [srb2kart socket to json translator](https://godoc.formulabun.club/pkg/go.formulabun.club/translator/).

## Technologies used

Everything is written in [Go](https://go.dev/), using the teachings form [gopl](http://www.gopl.io/). It is a simple language but draws a lot of power from that simplicity. [formulabun.club]() runs on a unix [VPS](https://en.wikipedia.org/wiki/Virtual_private_server) from [contabo](https://contabo.com/en/), a very good price/performance host. It runs [nginx](https://nginx.org/en/) with a certificate from [Let's Encrypt](https://letsencrypt.org/).

The deployment on the VPS is done using [docker](https://www.docker.com/) and [docker compose](https://docs.docker.com/compose/). This allows for a [microsevice](https://microservices.io/) like deployment which keeps each service simple, going hand in hand with the simplicity of the Go language. In order to avoid writing the boilerplate HTTP server and client code, the API's are defined in [openapi](https://www.openapis.org/), from which the server and client code is generated using the Go [openapi generators](https://openapi-generator.tech/). A process that happens only a few times per service, due to their simplicity and small size. Both the docker compose files and the api specifications are written in [YAML](https://yaml.org/).

## Running the services

Take a zip from the releases, extract it, open the template.env file and follow the constructions. Use `docker compose` to manage the services.

The `addons` service is meant to be used to link addons to the kart service. The repository of addons, from the `.env` file, are mapped to `/repo` and the kart server addons are mapped to `/addonsmain`. Link addons from the repo to the kart using `ls -s /repo/<addon> /addonsmain/<folder>/`. Use this service with `docker compose run`.

## Setting up a development environment

First of, **THIS IS NOT THE REPOSITORY TO DEVELOP IN.** This is only a collection of the packages, mostly serves the godoc site on [godoc.formulabun.club]() and contains build and release docker compose files. It is not really meant for development.

Instead, clone the packages (the submodules in this repository) you want to work on in `GOPATH/go/src/go.formulabun.club/`. 

If this is the only package you're working in you're already good to go. Pick an issue, make a branch and create a pull request when you're done.

If you're working in multiple package at once, clone them all and use a [replace directive](https://go.dev/doc/modules/gomod-ref#replace) in the necessary `go.mod` files. These will generally take the form of `go.formulabun.club/<package> => ../<package>`.

Because Go imports can't form cycles, you can complete the pull requests for each package in order. Do this without breaking other packages. Avoid writing breaking code or having to version a package. This may be a challenge but you should write code that's isn't prone to breaking changes anyway 😉. Think of the future and write wisely.

### Running the development environment

You can either always build and run services in docker or setup the environment yourself.

##### Running in docker (WIP don't trust this)

Update the context paths in the [docker-compose](https://github.com/formulabun/go.formulabun.club/blob/master/docker-compose.yml) file in this repository, then just `docker compose build` and `up` to build and run a service. 

##### Setting up the environment

This can be service dependent. See the docker and docker compose files as reference. Setting up your environment consists of 

* exporting environment variables
* setting up static hostname redirects
* creating folders

Each is OS dependent.

