package interfaces_internal

type IUser struct {
	*IAccount
	FirstName    string       `bson:"first_name"`
	MiddleName   string       `bson:"middle_name"`
	LastName     string       `bson:"last_name"`
	Birthday     string       `bson:"birthday"`
	Gender       string       `bson:"gender"`
	PersonalInfo IUserProfile `bson:"personal_info"`
	Experiences []IExperience `bson:"experiences"`
}
