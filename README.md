# hello-world
Hello world client and server application to be used for checking the network


#gRPC

## Server

Info:
```
hello-world grpc server -h
Start a gRPC hello server

Usage:
  hello-world grpc server [flags]

Flags:
  -h, --help   help for server

Global Flags:
      --cert string     authority to use when calling
      --key string      authority to use when calling
      --mtls            Set security to mTLS
      --port string     default port to use (default "5000")
      --rootCA string   authority to use when calling
      --tls             Set security to TLS
```


In order to start a gRPC server:
```bash
hello-world grpc server
```

To start server in TLS mode:
```bash
hello-world grpc server --tls --cert ./cert.crt --key ./cert.key
```

To start server in mTLS mode:
```bash
hello-world grpc server --mtls --cert ./cert.crt --key ./cert.key
```

## Client

Info:
```
Start a gRPC hello client

Usage:
  hello-world grpc client [flags]

Flags:
      --authority string   authority to use when calling
  -h, --help               help for client
      --timeout int        call timeout (default 30)

Global Flags:
      --cert string     authority to use when calling
      --key string      authority to use when calling
      --mtls            Set security to mTLS
      --rootCA string   authority to use when calling
      --tls             Set security to TLS
```

In order to call the gRPC server:
```bash
hello-world grpc client <address>:<port>
```

In order to call the gRPC server with TLS:
```bash
hello-world grpc client --tls --rootCA ./rootCA.crt <address>:<port>
```

In order to call the gRPC server with mTLS:
```bash
hello-world grpc client --mtls --rootCA ./rootCA.crt --cert ./clientcert.crt --key ./clientcert.key <address>:<port>
```

