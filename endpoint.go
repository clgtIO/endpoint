package endpoint

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Endpoint struct {
	Status string        `json:"status"`
	Data   interface{}   `json:"data"`
	Error  Error         `json:"error"`
}
