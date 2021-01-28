# build stage
FROM golang:latest
WORKDIR /app

# speed up
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

EXPOSE 8081

CMD ["/bin/sh", "/app/script/build.sh"]