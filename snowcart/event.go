package snowcart

type Event struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
	Value     int    `json:"value"`
	Timestamp int64  `json:"timestamp"`
}
