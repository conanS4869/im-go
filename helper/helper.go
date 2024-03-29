package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"github.com/zhangzhaojian/go-mongo/x/mongo/driver/uuid"
	"im-go/define"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(identity, email string) (string, error) {
	UserClaim := &UserClaims{
		Identity:         identity,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUUID() string {
	u, err := uuid.New()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", u)
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <87404816@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "87404816@qq.com", define.MailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	return err
}

func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}
