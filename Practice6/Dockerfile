FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod /app/
RUN go mod download
COPY main.go /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main main.go

FROM scratch
COPY --from=builder /app/main /app
EXPOSE 8080
CMD ["/app"]