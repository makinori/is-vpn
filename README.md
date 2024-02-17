# Is VPN

Docker image to determine if you're on a VPN network.

-   `/` to receive a html page
-   `/api` to receive JSON

Services available:

-   `mullvad`
-   `nordvpn`
-   `expressvpn`
-   `surfshark`

## Run

```bash
docker run -it --rm -p 8080:8080 -e SERVICE=mullvad ghcr.io/makidoll/is-vpn:latest
```

```yml
version: "3.6"
services:
    is-vpn:
        image: ghcr.io/makidoll/is-vpn:latest
        restart: always
        ports:
            - 8080:8080
        environment:
            - SERVICE=nordvpn
```

## Build and publish

```bash
docker build -t ghcr.io/makidoll/is-vpn:latest .
docker push ghcr.io/makidoll/is-vpn:latest
```
