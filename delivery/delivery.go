package delivery

type Delivery struct {
	MessageId string `json:"message_id"`
	Address   string `json:"address"`
	Delivered int64  `json:"delivered"`
}
