FROM golang:1.19-buster


RUN mkdir app

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["go", "run", "/app/cmd/botapi/main.go"]




#       FROM postgres:12.0

# COPY dockerDB/*.sql /docker-entrypoint-initdb.d/
# RUN chmod 755 /docker-entrypoint-initdb.d/*.sql