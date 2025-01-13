package entities

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type UrlShorten struct {
	Id        bson.ObjectID `bson:"_id" json:"id"`
	Token     string        `bson:"token" json:"token"`
	Url       string        `bson:"url" json:"url"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	ExpiredAt time.Time     `bson:"expiredAt" json:"expiredAt"`
}

func NewUrlShorten(token, url string, expiredAtDay int) *UrlShorten {
	return &UrlShorten{
		Token:     token,
		Url:       url,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().AddDate(0, 0, expiredAtDay),
	}
}

func EnsureUrlShortenIndex() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{"token", 1},
			},
			Options: options.Index().SetName("Ix_Asc_Token"),
		},
		{
			Keys: bson.D{
				{"url", 1},
			},
			Options: options.Index().SetName("Ix_Asc_Url"),
		},
	}
}
