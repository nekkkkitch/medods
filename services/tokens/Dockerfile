FROM nekkkkitch/docker AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o tokens ./services/tokens/cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tokens .
COPY ./services/tokens/cfg.yml .
CMD ["./tokens"]