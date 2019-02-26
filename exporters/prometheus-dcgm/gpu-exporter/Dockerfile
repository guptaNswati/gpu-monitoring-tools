FROM golang:1.11 as build

WORKDIR /go/src/gpu-exporter
COPY src .

RUN go install -v gpu-exporter

FROM debian:stretch-slim

COPY --from=build /go/bin/gpu-exporter /usr/bin/gpu-exporter

ENTRYPOINT ["gpu-exporter"]