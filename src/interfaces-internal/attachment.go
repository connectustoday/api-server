package interfaces_internal

type IAttachment struct {
	SchemaVersion uint32 `bson:"schema_version"`
	Type string `bson:"type"`
	URL string `bson:"url"`
	Description string `bson:"description"`
}