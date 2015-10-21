.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	proto \
	generate \
	docker-build-flights-dev \
	docker-build-flightsd-internal \
	docker-build-flightsd \
	launch

all: test

deps:
	go get -d -v ./...

updatedeps:
	go get -d -v -u -f ./...

testdeps:
	go get -d -v -t ./...

updatetestdeps:
	go get -d -v -t -u -f ./...

build: deps
	go build ./...

install: deps
	go install ./...

lint: testdeps
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go' | grep -v '\.pb.log\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

vet: testdeps
	go vet ./...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck ./...

pretest: lint

test: testdeps pretest
	go test ./...

clean:
	go clean -i ./...

proto:
	go get -v go.pedge.io/tools/protoc-all
	STRIP_PACKAGE_COMMENTS=1 protoc-all go.pedge.io/openflights

generate:
	go run cmd/gen-flights-csv-store/main.go flights generated.go

docker-build-flights-dev:
	docker build -t pedge/flights-dev -f Dockerfile.flights-dev .

docker-build-flightsd-internal: deps
	rm -rf _tmp
	mkdir -p _tmp
	go build \
		-a \
		-installsuffix netgo \
		-tags netgo \
		-ldflags '-w -linkmode external -extldflags "-static"' \
		-o _tmp/flightsd \
		cmd/flightsd/main.go
	docker build -t pedge/flightsd -f Dockerfile.flightsd .

docker-build-flightsd: docker-build-flights-dev
	docker run -v /var/run/docker.sock:/var/run/docker.sock pedge/flights-dev make docker-build-flightsd-internal

launch: docker-build-flightsd
	docker run -d -p 1747:1747 -p 8080:8080 pedge/flightsd
