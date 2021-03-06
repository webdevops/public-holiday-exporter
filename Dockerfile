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
FROM gcr.io/distroless/static
ENV LOG_JSON=1
COPY --from=build /go/src/github.com/webdevops/public-holiday-exporter/public-holiday-exporter /
USER 1000
ENTRYPOINT ["/public-holiday-exporter"]
