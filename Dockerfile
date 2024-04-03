FROM cgr.dev/chainguard/go:latest

ARG VERSION
ARG COMMIT

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -trimpath -ldflags "-s -w \
	-X main.AppVersion=${VERSION} \
	-X main.BuildCommit=${COMMIT}" \
	-o /usr/local/bin/unch .

ENTRYPOINT ["unch"]
