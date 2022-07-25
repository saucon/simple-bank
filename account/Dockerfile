FROM golang:1.18-alpine as builder

WORKDIR $GOPATH/src/github.com/Saucon/simple-bank/account


# COPY go.mod, go.sum and download the dependencies
COPY go.* ./

ENV GO111MODULE=on
ENV GOFLAGS=-mod=mod

# COPY All things inside the project and build
COPY . .
RUN go build $GOPATH/src/github.com/Saucon/simple-bank/account/cmd/accountapp
RUN ls -al

FROM alpine:latest

# Add curl
RUN apk --no-cache add curl
# SET TZ
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY --from=builder /go/src/github.com/Saucon/simple-bank/account/accountapp simple-bank/account/
COPY --from=builder /go/src/github.com/Saucon/simple-bank/account/.env ./

RUN ls -al /simple-bank/account

EXPOSE 9980
ENTRYPOINT [ "simple-bank/account/accountapp" ]
