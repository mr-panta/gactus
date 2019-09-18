package tcp

type Client interface {
	Send(input []byte) (output []byte, err error)
	Close() (err error)
}
