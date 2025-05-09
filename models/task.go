package models

type User struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Age   int    `json:"age,omitempty" bson:"age,omitempty"`
}

//var client *mongo.Client
//
//func New(mongo *mongo.Client) User {
//	client = mongo
//
//	return User{}
//}
