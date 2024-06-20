FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o usher-ifcb .

FROM debian:12-slim

COPY --from=build /app/usher-ifcb /usr/local/bin/usher-ifcb

ENTRYPOINT ["usher-ifcb"]
