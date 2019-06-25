package eventlog

type EventLog struct {
	Address string `json:"address"`
	Created int64  `json:"created"`
	Event   string `json:"event"`
	Message string `json:"message"`
}
