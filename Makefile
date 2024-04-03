APP      := unch
PKG      := ./pkg/unch/...
VERSION  := $(shell git describe --always --tags)
COMMIT   := $(shell git rev-parse HEAD | cut -c1-8)
LDFLAGS  := -s -w \
			-X main.AppVersion=${VERSION} \
			-X main.BuildCommit=${COMMIT}

vet:
	@go vet ./...

tidy:
	@go mod tidy

verify:
	@go mod verify

lint:
	@golangci-lint run ./...

test:
	@go test -v -race -count 1 -run "^Test" -cover ${PKG}

test-bench:
	@go test -race -count 1 -run "^$$" -bench . -cpu 4 -benchmem ${PKG}

test-fuzz:
	@go test -race -count 1 -run "^$$" -fuzz . -fuzztime 30s ${PKG}

test-all: test test-bench test-fuzz

ci: tidy verify vet test

cover: OUT := coverage.out
cover:
	@go test -v -race -count 1 -run "^Test" \
		-coverprofile=${OUT} \
		-covermode=atomic \
		-cover ${PKG}
	@go tool cover -func=${OUT}

build:
	@go build -v -trimpath -ldflags "${LDFLAGS}" -o ./bin/${UNCH} .

clean:
	@rm -rfv ./bin

clean-docker:
	@docker image rm ${APP}:latest ${APP}:${VERSION} -f

docker: clean-docker
docker:
	@docker build -t ${APP}:latest \
		--build-arg "VERSION=${VERSION}" \
		--build-arg "COMMIT=${COMMIT}" \
		--no-cache .
	@docker tag ${APP}:latest ${APP}:${VERSION}