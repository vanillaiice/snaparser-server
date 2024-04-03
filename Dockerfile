FROM golang:1.22.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY *.go ./

COPY cmd/snaparser_server/main.go ./cmd/snaparser_server/main.go

RUN go build -ldflags="-s -w" -o /snaparser_server ./cmd/snaparser_server/main.go

FROM scratch

WORKDIR /

COPY --from=build /snaparser_server /snaparser_server

EXPOSE 8888

ENTRYPOINT ["/snaparser_server"]

CMD ["--help"]
