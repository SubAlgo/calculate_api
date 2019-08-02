package cal

type calRequest struct {
	N1 int `json:"n1"`
	N2 int `json:"n2"`
}

type calResponse struct {
	Sum int `json:"sum"`
}

func cal(req calRequest) (res calResponse, err error) {
	res.Sum = req.N1 + req.N2
	return
}