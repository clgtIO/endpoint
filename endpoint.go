package endpoint

type Data interface{}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Endpoint struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
	Error  Error  `json:"error"`
}
