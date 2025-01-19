package entities

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type UrlToken struct {
	Id        bson.ObjectID `bson:"_id" json:"id"`
	Token     string        `json:"token"`
	IsUsed    bool          `json:"isUsed"`
	CreatedAt time.Time     `json:"createdAt"`
	UsedAt    time.Time     `json:"usedAt,omitempty"`
}

func NewUrlToken(token string) *UrlToken {
	return &UrlToken{
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
