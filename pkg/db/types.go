package db

type LoggerInterface interface {
	Fatal(...interface{})
}

type Entry struct {
	Text   string  `json:"text" bson:"text"`
	Number float64 `json:"number" bson:"number"`
	Found  bool    `json:"found" bson:"found"`
	Type   string  `json:"type" bson:"type"`
}
