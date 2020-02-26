FROM golang:1.14 as build

WORKDIR /go/src/github.com/webdevops/public-holiday-exporter

# Get deps (cached)
COPY ./go.mod /go/src/github.com/webdevops/public-holiday-exporter
COPY ./go.sum /go/src/github.com/webdevops/public-holiday-exporter
RUN go mod download

# Compile
COPY ./ /go/src/github.com/webdevops/public-holiday-exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /public-holiday-exporter \
    && chmod +x /public-holiday-exporter
RUN /public-holiday-exporter --help

#############################################
# FINAL IMAGE
#############################################
FROM gcr.io/distroless/static
COPY --from=build /public-holiday-exporter /
USER 1000
ENTRYPOINT ["/public-holiday-exporter"]
