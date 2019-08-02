package cal

import (
	"net/http"
	"github.com/subalgo/mod/internal/pkg/transport"
	)

var t = transport.HTTP{
	ErrorToStatusCode: errorToStatusCode,
	ErrorToMessage:    errorToMessage,
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/calculate", calHandler)
	return mux
}

func calHandler(w http.ResponseWriter, r *http.Request) {
	var req calRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := cal(req)
	t.EncodeResult(w, res, err)
}
