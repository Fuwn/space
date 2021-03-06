FROM golang:1.17.8-alpine3.14 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/space

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/space .

FROM alpine:3.15

RUN apk add ca-certificates

COPY --from=build_base /tmp/space/out/space /app/space

WORKDIR /app

EXPOSE 1965

CMD ["./space"]
