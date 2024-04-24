# RadioAnty [![Docker Image CI](https://github.com/Fisiu/radioanty/actions/workflows/docker-image.yml/badge.svg?branch=master)](https://github.com/Fisiu/radioanty/actions/workflows/docker-image.yml)

It seems that Antyradio has recently changed its streaming services. They are using some kind of load balancer now. Sending a request to the main stream results in adding a query param with a timestamp and then a redirection to a specific server instance.

Some streaming devices (like mine Revo SuperConnect Stereo) do not handle redirects properly, and this app solves above issue - it handles the redirect, adds a timestamp and returns the stream.

# Docker
Simply run a docker container with the following command:

`docker run -it --rm --name radioanty -p 8088:8088 fisiu82/radioanty:latest`

# Docker Compose

Create a `docker-compose.yml` with the following content:

```yaml
---
version: "3"

services:
  radioanty:
    image: fisiu82/radioanty:latest
    container_name: radioanty
    restart: unless-stopped
    ports:
      - 8088:8088/tcp
```

And then run `docker compose up -d` or `docker-compose up -d` depending on compose installation method.