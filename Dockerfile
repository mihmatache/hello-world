FROM alpine:3.7


COPY _bin/hello-world /usr/bin/hello-world

ENTRYPOINT [ "hello-world" ]