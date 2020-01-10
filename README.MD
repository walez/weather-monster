# Introduction

Domain object i.e key concepts e.g city, temperature etc are made stateless and are defined in the top level `core` package
When top level interface are modified `go generate` is handy to rebuild mocks struct

**Feature documentation**

Each feature should preferably have a `FEATURE.md` which
documents its behaviors, expectations, concepts, key
entities and entry points.

Here is the list of features:

- [weather](./weather/FEATURE.MD)

## Installation

- Make sure that [Golang](https://golang.org/) is installed, the mininum required is currently 1.12.
- To run migrations ensure [golang-migrate](https://github.com/golang-migrate/migrate) is installed, migration guide
- Migration [guide](./datastore/postgres/migrations/readme.md)
- Create `.env` file using `.env.example`

## Starting App

- Start the app with:

```shell
./run-local-app.sh
```

You should see something similiar:

```shell
INFO[0000] Starting the Weather Service!
INFO[0000] COMMIT:
INFO[0000] BRANCH:
INFO[0000] Connecting to postgres
INFO[0000] Connected to postgres
INFO[0000] Registering events manager
```

## Testing

There are two test coverage

- Datastore [test](./datastore/postgres/readme.md): for testing database query and insertation
- Routes test: end to end test for rest endpoint

## Links & Resources

- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://about.sourcegraph.com/go/idiomatic-go
- https://github.com/golang/go/wiki/Modules
