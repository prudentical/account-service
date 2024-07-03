FROM golang:1.22.2-alpine3.19 as build
WORKDIR /app

ENV GOPROXY https://goproxy.io

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY config.yml .

RUN CGO_ENABLED=0 GOOS=linux go build -o account-service ./cmd/application/


FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=build /app/internal/database/migrations ./internal/database/migrations
COPY --from=build /app/config.yml .
COPY --from=build /app/account-service .

EXPOSE 8001

CMD ["./account-service"]