# Is VPN

https://hub.docker.com/r/makitsune/is-vpn

Docker container to determine if you're on a VPN network.

-   `/` to receive a html page
-   `/api` to receive JSON

Services available:

-   `nordvpn`
-   `expressvpn`

## Run

```bash
docker run -it --rm -p 8080:8080 -e SERVICE=nordvpn makitsune/is-vpn
```

```yml
version: "3.6"
services:
    is-vpn:
        image: makitsune/is-vpn
        restart: always
        ports:
            - 8080:8080
        environment:
            - SERVICE=nordvpn
```

## Build

```bash
docker build -t makitsune/is-vpn .
```
