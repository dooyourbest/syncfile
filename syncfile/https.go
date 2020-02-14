package syncfile

const (
	CA_ROOT = "C:\\Users\\Administrator\\go\\src\\github.com\\dooyourbest\\syncfile\\syncfile\\ca\\"
	CA_PATH = CA_ROOT+"ca.key"
	CA_CRT = CA_ROOT+"ca.crt"

	SERVER_KEY = CA_ROOT+"server.key"
	SERVRT_CRT = CA_ROOT+"server.crt"

	CLIENT_KEY = CA_ROOT+"client.key"
	CLIENT_CRT = CA_ROOT+"client.crt"
)

type httpClient struct {
	useHttps bool
	caPath string
	clientKey string
	clientCrt string
}
type httpServer struct {

}

func checkssl(){

}
