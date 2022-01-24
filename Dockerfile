FROM golang:alpine3.15 as builder
# RUN apk add make
# RUN apk add make git gcc musl-dev
COPY . /opt/goscript
WORKDIR /opt/goscript
RUN mkdir -p ./bin;\
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64;\
        go build -o bin/goscript .;
# RUN ["make","build_alpine"]

FROM alpine:3.15.0
# RUN apk update && apk upgrade
RUN mkdir -p /usr/local/goscript
COPY --from=builder /opt/goscript/bin/goscript /usr/local/goscript
RUN ln -s /usr/local/goscript/goscript /usr/bin/goscript
STOPSIGNAL SIGTERM
EXPOSE 3034
EXPOSE 3032
WORKDIR /usr/local/goscript
RUN chmod +x /usr/local/goscript/goscript
CMD [ "/usr/local/goscript/goscript" ]