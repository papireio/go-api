FROM golang:1.21

WORKDIR /usr/src
COPY . .
RUN make vendor \
    && make build \
    && mkdir -p /usr/app \
    && cp ./target/go-api /usr/app \
    && cp -r ./swagger-ui /usr/app

WORKDIR /usr/app
RUN rm -rf /usr/src

CMD ["./go-api"]
