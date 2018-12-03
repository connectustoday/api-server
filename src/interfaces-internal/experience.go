package interfaces_internal

type IExperience struct {
	SchemaVersion uint32 `bson:"schema_version"`
	Location IAddress `bson:"address"`
	Name string `bson:"name"`
	Organization string `bson:"organization"`
	Opportunity string `bson:"opportunity"`
	Description string `bson:"description"`
	When string `bson:"when"`
	IsVerified bool `bson:"is_verified"`
	EmailVerify bool `bson:"email_verify"`
	CreatedAt int64 `bson:"created_at"`
	Hours int64 `bson:"hours"`
	EmailJWT string `bson:"emailjwt"`
}