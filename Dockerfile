FROM golang:1.23.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -extldflags '-static'" -o phone_validator cmd/main.go

FROM gcr.io/distroless/static-debian12:latest-amd64 AS release-stage

WORKDIR /

COPY --from=build-stage /app/phone_validator /phone_validator

EXPOSE 7777

USER nonroot:nonroot

ENTRYPOINT [ "./phone_validator" ]
