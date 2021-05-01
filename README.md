# Is ExpressVPN

Docker container to determine if you're on the ExpressVPN network.

-   `/` to receive a html page
-   `/api` to receive JSON

## Run

```bash
docker run -it --rm -p 3000:3000 makitsune/is-expressvpn
```

```yml
version: "3.6"
services:
    is-expressvpn:
        image: makitsune/is-expressvpn
        restart: always
        ports:
            - 3000:3000
```

## Build

```bash
docker build -t makitsune/is-expressvpn .
# or
build.bat
```
