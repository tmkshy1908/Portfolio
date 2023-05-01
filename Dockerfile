FROM golang:1.16.3-buster

WORKDIR /cmd/botapi

COPY /cmd/botapi /app

EXPOSE 8080

CMD ["go", "run", "cmd/botapi/main.go"]