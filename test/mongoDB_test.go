package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	NickName  string `bson:"nickName"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

type UserRoom struct {
	UserIdentity    string `bson:"user_identity"`
	RoomIdentity    string `bson:"room_identity"`
	MessageIdentity string `bson:"message_identity"`
	RoomType        int    `bson:"room_type"` //0-删除好友 1-私人聊天 2-群聊
	CreatedAt       int64  `bson:"created_at"`
	UpdatedAt       int64  `bson:"updated_at"`
}

// 测试MongoDB链接
func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username:    "admin",
		Password:    "admin",
		PasswordSet: false,
	}).ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		t.Fatal(err)
	}
	//连接数据库
	db := client.Database("im")

	cur := db.Collection("user_basic")

	ub := new(UserBasic)
	err = cur.FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ub)
}

// 测试MongoDB链接
func TestFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username:    "admin",
		Password:    "admin",
		PasswordSet: false,
	}).ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		t.Fatal(err)
	}
	//连接数据库
	db := client.Database("im")

	cur, err := db.Collection("user_room").Find(context.Background(), bson.D{})
	urs := make([]*UserRoom, 0)
	for cur.Next(context.Background()) {
		ur := new(UserRoom)
		err := cur.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}
	for _, ur := range urs {
		fmt.Println(ur)
	}
}
