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
	docker-build-openflights-dev \
	docker-build-openflightsd-internal \
	docker-build-openflightsd \
	launch \
	launch-local

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
	go run cmd/gen-openflights-csv-store/main.go openflights generated.go

docker-build-openflights-dev:
	docker build -t pedge/openflights-dev -f Dockerfile.openflights-dev .

docker-build-openflightsd-internal: deps
	rm -rf _tmp
	mkdir -p _tmp
	go build \
		-a \
		-installsuffix netgo \
		-tags netgo \
		-ldflags '-w -linkmode external -extldflags "-static"' \
		-o _tmp/openflightsd \
		cmd/openflightsd/main.go
	docker build -t pedge/openflightsd -f Dockerfile.openflightsd .

docker-build-openflightsd: docker-build-openflights-dev
	docker run -v /var/run/docker.sock:/var/run/docker.sock pedge/openflights-dev make docker-build-openflightsd-internal

launch: docker-build-openflightsd
	docker run -d -p 1747:1747 -p 8080:8080 pedge/openflightsd

launch-local: deps
	go run cmd/openflightsd/main.go
