# syntax=docker/dockerfile:1.4

FROM golang:1.21-alpine AS build-dev
WORKDIR /go/src/app
COPY --link go.mod ./
RUN go version && go mod download
COPY --link . .
RUN CGO_ENABLED=0 go install -buildvcs=false -trimpath -ldflags '-w -s'
FROM scratch
COPY --link --from=build-dev /go/bin/go-app /go/bin/go-app
COPY --from=build-dev /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD ["/go/bin/go-app"]
