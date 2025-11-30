FROM golang:1.25.4 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./todo-back


FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=builder /app/todo-back .
EXPOSE 3000
ENTRYPOINT ["./todo-back"]

# docker build --tag todo-back:1.0.0 .