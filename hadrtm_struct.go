package  packagertm   

import "go.mongodb.org/mongo-driver/bson/primitive"

type grapic struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Pendapatan    string             `bson:"pendapatan,omitempty" json:"pendapatan,omitempty"`
	Laba string                      `bson:"laba ,omitempty" json:"laba,omitempty"`
	Keuntungan  string               `bson:"keuntungan,omitempty" json:"keuntungan,omitempty"`
	Kerugian  string                 `bson:"kerugian,omitempty" json:"kerugian,omitempty"`
}