FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o fast_sns_service_proxy .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/fast_sns_service_proxy /app/
WORKDIR /app
CMD ["./fast_sns_service_proxy"]
