FROM alpine:3.7

WORKDIR app
COPY ./ ./

ENTRYPOINT ["/app/monitor"]