# m3u-proxy
A simple m3u proxy server that can be used to proxy m3u files to clients.
It allows simple rewriting of the m3u-urls by configuration.

# Installation
## With prebuilt binary
A prebuilt ubuntu binary can be downloaded from the [releases](https://github.com/iptv-blog/m3u-proxy/releases) page.

## With Docker
```bash
docker run -d -p 8080:8080 iptvblog/m3u-proxy
``` 

## With Go
With local go installation you can build it from source
```bash
go install github.com/iptv-blog/m3u-proxy
```

# Usage
Adapt the `config.yml` to your needs and start the server with `m3u-proxy`.