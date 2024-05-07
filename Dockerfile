FROM alpine:3.19
COPY m3u-proxy /m3u-proxy
EXPOSE 8080
ENTRYPOINT ["/m3u-proxy"]