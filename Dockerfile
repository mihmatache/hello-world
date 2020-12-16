FROM alpine:3.7


COPY _build/hello /usr/bin/hello

ENV PORT=5000

CMD shop grpc server --port ${PORT}