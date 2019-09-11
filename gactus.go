package gactus

type Request struct {
	Path        string
	Method      string
	ContentType string
	Body        []byte
}

type Response struct{}
