FROM golang:1.25.4 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./todo-back


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/todo-back .
EXPOSE 8080
ENTRYPOINT ["./todo-back"]

# docker build --tag todo-back .