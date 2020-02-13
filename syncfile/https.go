package syncfile

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	CA_PATH = "./ca/ca.key"
	CA_CRT = "./ca/ca.crt"

	SERVER_KEY = "./ca/server.key"
	SERVRT_CRT = "./ca/server.crt"

	CLIENT_KEY = "./ca/client.key"
	CLIENT_CRT = "./ca/client.crt"
)

type httpClient struct {
	useHttps bool
	caPath string
	clientKey string
	clientCrt string
}
type httpServer struct {

}
func getHttpsClient()(http.Client)  {
		pool := x509.NewCertPool()
		caCrt, err := ioutil.ReadFile(CA_CRT)
		if err != nil {
			fmt.Println("ReadFile err:", err)
		}
		pool.AppendCertsFromPEM(caCrt)
		cliCrt, err := tls.LoadX509KeyPair(CLIENT_CRT,CLIENT_KEY)
		if err != nil {
			fmt.Println("Loadx509keypair err:", err)
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
			},
		}
		return http.Client{Transport: tr}
}
func checkssl(){
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(CA_CRT)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	s := &http.Server{
		Addr:    ":"+REMOTE_PORT,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err=s.ListenAndServeTLS(SERVRT_CRT,SERVER_KEY)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
