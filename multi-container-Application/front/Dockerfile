FROM node:22.bullseye AS builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

FROM busybox:1.30 AS runner
WORKDIR /app
COPY --from=builder /app/dist .
EXPOSE 80
CMD ["busybox", "httpd", "-f", "-v", "-p", "80"]

# docker build --tag todo-front:1.0.0 .
