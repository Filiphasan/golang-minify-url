package entities

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type UrlToken struct {
	Id        bson.ObjectID `bson:"_id" json:"id"`
	Token     string        `bson:"token" json:"token"`
	IsUsed    bool          `bson:"isUsed" json:"isUsed"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UsedAt    time.Time     `bson:"usedAt" json:"usedAt,omitempty"`
}

func NewUrlToken(token string) *UrlToken {
	return &UrlToken{
		Id:        bson.NewObjectID(),
		Token:     token,
		IsUsed:    false,
		CreatedAt: time.Now(),
		UsedAt:    time.Time{},
	}
}

func EnsureUrlTokenIndex() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"isUsed", 1},
				{"createdAt", 1},
			},
			Options: options.Index().SetName("Ix_Asc_IsUsed_Asc_CreatedAt"),
		},
		{
			Keys: bson.D{
				{"token", 1},
			},
			Options: options.Index().SetName("Ix_Asc_Token").SetUnique(true),
		},
	}
}
