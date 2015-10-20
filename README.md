[![CircleCI](https://circleci.com/gh/peter-edge/go-flights/tree/master.png)](https://circleci.com/gh/peter-edge/go-flights/tree/master)
[![GoDoc](http://img.shields.io/badge/api-godoc-blue.svg)](https://godoc.org/go.pedge.io/flights)
[![MIT License](http://img.shields.io/badge/license-mit-blue.svg)](https://github.com/peter-edge/go-flights/blob/master/LICENSE)

### Please donate to OpenFlights!

If you do use this, I ask you to donate to OpenFlights, the source for all the data
in here as of now, at http://openflights.org/donate. Seriously, if you can afford it, the OpenFlights
team is responsible for putting all this data together and maintaining it, and we owe it to them
to support their work.

## Introduction

Note the custom import path!

```go
import (
  "go.pedge.io/flights"
)
```

Flights is a package that aims to eventually expose APIs for all kinds of flights data. At the moment,
it exposes the data from http://openflights.org/data.html, available within https://github.com/jpatokal/openflights/tree/master/data.

Flights uses [protobuf](https://developers.google.com/protocol-buffers/docs/proto3) and [gRPC](http://www.grpc.io) to auto-generate
a protobuf/grpc API stubs, and a HTTP/JSON API. See [flights.proto](flights.proto) for the API definition. The HTTP endpoints
should be relatively straightfoward.

The binary [flightsd](cmd/flightsd) is a server binary that hosts the API. `make install` will install this, which you can
then run with `${GOPATH}/bin/flightsd`, or `make launch` will build a Docker image that is ~13MB as of now, and launch
this Docker image with the default ports set. Then, you can `curl http://0.0.0.0:8080/airports/code/sfo` as a quick check.

The flights package for golang adds some golang-specific additional functionality around the generated protocol buffers code.
See [flights.go](flights.go) for publically-exposed structures.

The binaries [gen-flights-csv-store](cmd/gen-flights-csv-store) and [gen-flights-id-store](cmd/gen-flights-id-store) will generate
the data for you in either a `CSVStore` or `IDStore`. Note that this is how [generated.go](generated.go) is generated.

## Future Work

This is just a start to what this will be. Please contact me if you want to help out!
