FROM alpine:3.7


COPY _bin/hello-world /usr/bin/hello-world

ENV PORT=5000

CMD hello-world grpc server --port ${PORT}