package interfaces_internal

type IUserProfile struct {
	SchemaVersion    uint32   `bson:"schema_version"`
	Interests        []string `bson:"interests"`
	Biography        string   `bson:"biography"`
	Education        string   `bson:"education"` // TODO
	Quote            string   `bson:"quote"`
	CurrentResidence string   `bson:"current_residence"`
	Certifications   string   `bson:"certifications"` // TODO
}
