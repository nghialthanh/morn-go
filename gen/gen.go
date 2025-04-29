package gen

type Generator struct {
	ID    string `bson:"_id"`
	Value int64  `bson:"value"`
}
