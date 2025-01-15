FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /main .
COPY path/to/your/largefile.csv ./largefile.csv
ENTRYPOINT ["./main"]
