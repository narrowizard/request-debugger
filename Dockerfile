FROM golang:1.12

WORKDIR /go/src/github.com/narrowizard/request-debugger

COPY . .

ENV GOPATH /go

ARG ROOT=narrowizard/request-debugger

ARG TARGET=request-debugger

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64                      \
	go build -i -v -o /tmp/${TARGET}

FROM alpine

ARG TARGET=request-debugger

COPY --from=0 /tmp/${TARGET} /${TARGET}

RUN ln -s /${TARGET} /entrypoint

EXPOSE 80

CMD ["/entrypoint"]
