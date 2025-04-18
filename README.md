# Is VPN

<img align="right" height="200" src="https://github.com/user-attachments/assets/33f74534-07ff-4798-be6a-2913ff6d2031" />

Docker image to determine if you're on a VPN network.

-   `/` to receive a html page
-   `/api` to receive JSON

Services available:

_\*some may be broken_

-   `expressvpn`\*
-   `mullvad`\*
-   `nordvpn`\*
-   `privateinternetaccess` or `pia`
-   `surfshark`\*

## Run

```bash
podman run -it --rm -p 8080:8080 -e SERVICE=pia ghcr.io/makinori/is-vpn:latest
```

```yml
services:
    is-vpn:
        image: ghcr.io/makinori/is-vpn:latest
        restart: always
        # network: service:vpn
        ports:
            - 8080:8080
        environment:
            - SERVICE=pia
```

## Build and publish

```bash
podman build -t ghcr.io/makinori/is-vpn:latest .
podman push ghcr.io/makinori/is-vpn:latest
```
