FROM golang:1.14-alpine AS build

RUN apk update && apk upgrade && \
    apk add --no-cache git

# Add Maintainer Info
LABEL maintainer="Aayush Joglekar <aayushjog@gmail.com>"

WORKDIR /tmp/app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux go build -o ./out/bunkr .

FROM alpine:latest

RUN apk add ca-certificates

COPY --from=build /tmp/app/out/bunkr /app/bunkr
COPY --from=build /tmp/app/.env /app/.env

WORKDIR "/app"

EXPOSE 5000

CMD ["./bunkr"]