package certificates

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"google.golang.org/grpc/credentials"
)

func SetRootTLSCert(config *tls.Config, rootCA []byte) error {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCA) {
		return fmt.Errorf("could not add certificate to cert pool")
	}
	config.RootCAs = certPool
	return nil
}

func SetTLSCertificates(conf *tls.Config, cert string, key string) error {
	fmt.Println(len(cert))
	fmt.Println(len(key))
	if len(cert) == 0 || len(key) == 0 {
		return fmt.Errorf("Please proivide a client certificate or key")
	}
	clientCertPair, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}
	conf.Certificates = []tls.Certificate{clientCertPair}
	return nil
}

func MakeCredentials(conf *tls.Config) credentials.TransportCredentials {
	return credentials.NewTLS(conf)
}
