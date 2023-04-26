package result

type Result struct {
	Code     Code        `json:"code"`
	Msg      string      `json:"message"`
	Response interface{} `json:"data"`
}
