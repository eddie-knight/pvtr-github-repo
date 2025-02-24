FROM golang:1.23.4-alpine3.21 AS build
RUN apk add --no-cache make git
WORKDIR /app
RUN git clone https://github.com/privateerproj/privateer.git .
RUN go mod tidy && \
    make go-build

FROM golang:1.23.4-alpine3.21 AS plugin
RUN apk add --no-cache make git
WORKDIR /plugin
COPY . .
RUN go mod tidy && \
    make build

FROM golang:1.23.4-alpine3.21
ARG log_level="debug"
ENV LOG_LEVEL=$log_level
RUN apk add --no-cache make git && \
    mkdir -p /.privateer/bin
WORKDIR /.privateer/bin
COPY --from=build /app/privateer .
COPY --from=plugin /plugin/pvtr-github-repo .

CMD ["./privateer", "run", "--binaries-path", ".", "--config", "../config.yml", "--loglevel", "${LOG_LEVEL}"]
