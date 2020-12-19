FROM golang:latest AS builder

ENV GOPROXY=https://goproxy.io,direct
ENV CGO_ENABLED=0

WORKDIR /root/
COPY . ./

WORKDIR /root/cmd/gateway
RUN go build -o gateway

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /root/cmd/gateway/gateway .
COPY cmd/gateway/application.toml .

EXPOSE 8001 

CMD [ "./gateway", "-c", "application.toml" ]
