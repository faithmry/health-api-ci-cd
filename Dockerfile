FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /health-api

FROM alpine:latest

# Install timezone data
RUN apk add --no-cache tzdata

WORKDIR /root/

COPY --from=build /health-api .

# Set the timezone
ENV TZ=Asia/Jakarta

EXPOSE 8080

CMD ["./health-api"]
