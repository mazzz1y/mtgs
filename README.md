# mtgs

[![Go Report Card](https://goreportcard.com/badge/github.com/dmirubtsov/mtgs)

Fork from [mtg](https://github.com/9seconds/mtg)

Bullshit MTPROTO proxy for Telegram with per-user tokens and RestAPI for that. Using [Consul](https://consul.io) as KV storage

## TODO

* Add docker-compose
* Add Helm chart for k8s

## API Documentation

Please see [API.md](API.md)

## Environment variables


| Environment variable          | Corresponding flags           | Default value                     | Description                                                                                                                                                                                                                                                                |
|-------------------------------|-------------------------------|-----------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `MTGS_DEBUG`                   | `-d`, `--debug`               | `false`                           | Run in debug mode. Usually, you need to run in this mode  only if you develop this tool or its maintainer is asking you to provide  logs with such verbosity.                                                                                                              |
| `MTGS_VERBOSE`                 | `-v`, `--verbose`             | `false`                           | Run in verbose mode. This is way less chatty than debug mode.                                                                                                                                                                                                              |
| `MTGS_IP`                      | `-b`, `--bind-ip`             | `127.0.0.1`                       | Which IP should we bind to. As usual, `0.0.0.0` means that we want to listen on all interfaces. Also, 4 zeroes will bind to both IPv4 and IPv6.                                                                                                                            |
| `MTGS_PORT`                    | `-P`, `--bind-port`           | `3128`                            | Which port should we bind to (listen on).                                                                                                                                                                                                                                  |
| `MTGS_IPV4`                    | `-4`, `--public-ipv4`         | [Autodetect](https://ifconfig.co) | IPv4 address of this proxy. This is required if you NAT your proxy or run it in a docker container. In that case, you absolutely need to specify public IPv4 address of the proxy, otherwise either URLs will be broken or proxy could not access Telegram middle proxies. |
| `MTGS_IPV4_PORT`               | `--public-ipv4-port`          | Value of `--bind-port`            | Which port should be public of IPv4 interface. This affects only generated links and should be changed only if you NAT your proxy or run it in a docker container.                                                                                                         |
| `MTGS_IPV6`                    | `-6`, `--public-ipv6`         | [Autodetect](https://ifconfig.co) | IPv6 address of this proxy. This is required if you NAT your proxy or run it in a docker container. In that case, you absolutely need to specify public IPv6 address of the proxy, otherwise either URLs will be broken or proxy could not access Telegram middle proxies. |
| `MTGS_IPV6_PORT`               | `--public-ipv6-port`          | Value of `--bind-port`            | Which port should be public of IPv6 interface. This affects only generated links and should be changed only if you NAT your proxy or run it in a docker container.                                                                                                         |
| `MTGS_BUFFER_WRITE`            | `-w`, `--write-buffer`        | `65536`                           | The size of TCP write buffer in bytes. Write buffer is the buffer for messages which are going from client to Telegram.                                                                                                                                                    |
| `MTGS_BUFFER_READ`             | `-r`, `--read-buffer`         | `131072`                          | The size of TCP read buffer in bytes. Read buffer is the buffer for messages from Telegram to client.                                                                                                                                                                      |
| `MTGS_SECURE_ONLY`             | `-s`, `--secure-only`         | `false`                           | Support only clients with secure mode (i.e only clients with dd-secrets).                                                                                                                                                                                                  |
| `MTGS_ANTIREPLAY_MAXSIZE`      | `--anti-replay-max-size`      | `128`                             | Max size of antireplay cache in megabytes.                                                                                                                                                                                                                                 |
| `MTGS_ANTIREPLAY_EVICTIONTIME` | `--anti-replay-eviction-time` | `168h`                            | Eviction time for antireplay cache entries.                                                                                                                                                                                                                                |
| `MTGS_API_PORT`                | `-p`, `--api-port`            | `8080`                            | Which port should be public for Rest API.                                                                                                                                                                                                                                  |
| `MTGS_API_PATH`                | `--api-path`                  | `/mtg`                            | Which basepath should be use for Rest API (It useful when API behind proxy).                                                                                                                                                                                               |
| `MTGS_CONSUL_HOST`             | `--consul-host`               | `127.0.0.1`                       | Address of Consul.                                                                                                                                                                                                                                                         |
| `MTGS_CONSUL_PORT`             | `--consul-port`               | `8050`                            | Port of Consul.                                                                                                                                                                                                                                                            |
| `MTGS_API_TOKEN`               | `-t`, `--api-token`           | ``                                | Which token use for API authorization. Should be provided in `Authorization` header. Auth disabled if this value is empty                                                                                                                                                 | 
