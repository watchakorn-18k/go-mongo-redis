package entities

type NewUserBody struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type UserDataFormat struct {
	ID        string `json:"id" bson:"id,omitempty"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Gender    string `json:"gender" bson:"gender"`
	IpAddress string `json:"ip_address" bson:"ip_address"`
}
