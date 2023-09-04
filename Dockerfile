FROM golang:1.19-buster

RUN mkdir app

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["go", "run", "/app/cmd/botapi/main.go"]

FROM envoyproxy/envoy:v1.18.3
COPY envoy/envoy.yaml /app/etc/envoy/envoy.yaml