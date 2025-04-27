package handler

type UserHandlerResponse struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}
