package web

type WebReponse struct {
	Code   int16       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
