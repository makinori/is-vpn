# Is VPN

-   https://github.com/makitsune/is-vpn
-   https://hub.docker.com/r/makitsune/is-vpn

Docker image to determine if you're on a VPN network.

-   `/` to receive a html page
-   `/api` to receive JSON

Services available:

-   `mullvad`
-   `nordvpn`
-   `expressvpn`

## Run

```bash
docker run -it --rm -p 8080:8080 -e SERVICE=mullvad makitsune/is-vpn
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
