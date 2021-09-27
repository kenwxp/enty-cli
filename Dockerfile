# Build HTC in a stock Go builder container
FROM golang:1.15-alpine as builder

ADD . /htc-backend
RUN cd /htc-backend && go build -o htc-backend cmd/backend/main.go

# Pull HTC into a second stage deploy alpine container
FROM alpine:latest

COPY --from=builder /htc-backend/htc-backend /usr/local/bin/

EXPOSE 8848
ENTRYPOINT ["htc-backend"]