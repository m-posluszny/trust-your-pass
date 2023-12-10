package passwords

import "go.mongodb.org/mongo-driver/bson/primitive"

type PasswordDto struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Preconditions []PreconditionDto  `json:"preconditions" bson:"preconditions,omitempty"`
	Strength      int                `json:"strength" bson:"strength"`
	IsProcessed   bool               `json:"isProcessed" bson:"isProcessed"`
}

type PreconditionDto struct {
	ConditionName string `json:"condition"`
	IsSatisfied   bool   `json:"isSatisfied"`
}

type PostResponseDto struct {
	Id            interface{}       `json:"id" bson:"id"`
	Preconditions []PreconditionDto `json:"preconditions" bson:"preconditions"`
	Strength      int               `json:"strength" bson:"strength"`
	IsProcessed   bool              `json:"isProcessed" bson:"isProcessed"`
}
