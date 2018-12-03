package interfaces_internal

type INotification struct {
	ID int64 `bson:"id"`
	CreatedAt int64 `bson:"created_at"`
	Type string `bson:"type"`
	Content string `bson:"content"`
	Account string `bson:"account"`
}
