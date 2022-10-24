FROM golang:1.19.0-alpine as builder

ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY . .

RUN go mod download && go mod verify
RUN go build  -o /out/customer ./customer && \
    go build  -o /out/flight ./flight && \
    go build  -o /out/booking ./booking

################################

FROM alpine:latest

RUN apk update && rm /var/cache/apk/*
RUN apk --no-cache add ca-certificates libressl-dev libffi-dev

WORKDIR /app

COPY --from=builder /out/* .
COPY ./entrypoint.sh .
RUN chmod -R +x /app/

ENTRYPOINT ["/app/entrypoint.sh"]
