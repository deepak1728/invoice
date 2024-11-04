package models

type Invoice struct {
	FirstName string `bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName  string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Gender    string `bson:"gender,omitempty" json:"gender,omitempty"`
	Email     string `bson:"email,omitempty" json:"email,omitempty"`
	Password  string `bson:"password,omitempty" json:"password,omitempty"`
}

type ErrorMsg struct {
	Message string `json:"message"`
}
