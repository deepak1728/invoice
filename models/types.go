package models

type Invoice struct {
	FirstName string `json:"firstname,omitempty", bson:"_firstname,omitempty"`
	LastName  string `json:"lastname,omitempty", bson:"_lasttname,omitempty"`
	Gender    string `json:"gender,omitempty", bson:"_gender,omitempty"`
	Email     string `json:"email,omitempty", bson:"_email,omitempty"`
	Password  string `json:"password,omitempty", bson:"password,omitempty"`
}
