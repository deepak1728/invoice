FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o main .
RUN ls -la /app

COPY .env .env

FROM alpine:latest
RUN apk --no-cache add libc6-compat

WORKDIR /root/


COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

CMD [ "./main" ]