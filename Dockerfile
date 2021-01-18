FROM golang:1.15 as build

WORKDIR /go/src/github.com/webdevops/public-holiday-exporter

# Get deps (cached)
COPY ./go.mod /go/src/github.com/webdevops/public-holiday-exporter
COPY ./go.sum /go/src/github.com/webdevops/public-holiday-exporter
COPY ./Makefile /go/src/github.com/webdevops/public-holiday-exporter
RUN make dependencies

# Compile
COPY ./ /go/src/github.com/webdevops/public-holiday-exporter
RUN make test
RUN make lint
RUN make build
RUN ./public-holiday-exporter --help

#############################################
# FINAL IMAGE
#############################################
FROM alpine:3.12
ENV LOG_JSON=1
RUN apk add bash tzdata
COPY ./docker/entrypoint.sh /entrypoint.sh
RUN mkdir /app 
COPY --from=build /go/src/github.com/webdevops/public-holiday-exporter/public-holiday-exporter /app/
RUN chown -R 1000:1000 /app
USER 1000
ENTRYPOINT /entrypoint.sh
EXPOSE 8080
