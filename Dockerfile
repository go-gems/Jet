FROM golang:1.16-alpine as BUILD
RUN apk --update upgrade \
    && apk --no-cache --no-progress add git bash gcc curl tar ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*
WORKDIR /go/src/github.com/go-gems/jet

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download

COPY . /go/src/github.com/go-gems/jet
RUN go build -o jet .

FROM alpine as RUN
COPY --from=BUILD /go/src/github.com/go-gems/jet/jet /usr/local/bin/jet
RUN chmod 777 /usr/local/bin/jet
EXPOSE 8000
ENTRYPOINT ["jet"]


