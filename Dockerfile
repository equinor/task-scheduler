#build the scheduler
FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scheduler .

#run the scheduler
FROM alpine
COPY --from=builder /app/scheduler /scheduler
USER 1000
ENTRYPOINT ["/scheduler"]
