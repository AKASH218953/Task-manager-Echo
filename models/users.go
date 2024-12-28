package models

type Users struct {
	ID        string `bson:"_id"`
	FristName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	UserName  string `bson:"username"`
	Password  string `bson:"password"`
	Role      string `bson:"role"`
}
