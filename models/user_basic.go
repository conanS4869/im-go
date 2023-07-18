package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func JudgeUserIsFriend(userIdentity1, userIdentity2 string) bool {
	// 查询 userIdentity1 单聊房间列表
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return false
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}
	// 获取关联 userIdentity2 单聊房间个数
	cnt, err := Mongo.Collection(UserRoom{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"user_identity": userIdentity2, "room_type": 1, "room_identity": bson.M{"$in": roomIdentities}})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}

	return false
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}}).
		Decode(ub)
	return ub, err
}
func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"identity", identity}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicCountByEmail(email string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}
func InsertOneUserBasic(ub *UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).InsertOne(context.Background(), ub)
	return err
}

func GetUserBasicCountByAccount(account string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"account", account}})
}
