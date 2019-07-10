package methodmux

import (
	"errors"
)

var (
	errFound                   = errors.New("methodmux:Found")
	errNilHandler              = errors.New("methodmux: Nil Handler")
	errNilReceiver             = errors.New("methodmux: Nil Receiver")
	errOptionsMethodNotAllowed = errors.New("methodmux: HTTP OPTIONS Method Not Allowed")
)
