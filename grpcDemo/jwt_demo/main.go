package jwtUtils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTWithKey(jwtKey []byte) string {
	/*
		key：该参数是用于签名 token 的密钥。密钥的类型取决于使用的签名算法。
		例如，如果使用 HMAC 算法（如 HS256、HS384 等），key 应该是一个对称密钥（通常是 []byte 类型的密钥）。
		如果使用 RSA 或 ECDSA 签名算法（如 RS256、ES256），key 应该是一个私钥 *rsa.PrivateKey 或 *ecdsa.PrivateKey
	*/
	method := jwt.SigningMethodHS512
	claims := jwt.MapClaims{}
	claims["name"] = "hydra"
	claims["exp"] = time.Now().Add(time.Second * 10).UnixMilli()
	token := jwt.NewWithClaims(method, claims, nil)
	jwtString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("jwtString is %v\n", jwtString)
		return jwtString
	}
}
func ParseJWTStringWithOfficialClaims() {
	jwtKey := make([]byte, 32) // 生成32字节（256位）的密钥

	if _, err := rand.Read(jwtKey); err != nil { //rand.Read函数可以用来生成随机字节序列
		panic(err)
	}
	jwtString := GenerateJWTWithKey(jwtKey)
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }, jwt.WithExpirationRequired()) //第二个参数是一个回调函数，返回用于验证 JWT 签名的密钥。该函数签名为 func(*Token) (interface{}, error)
	if err != nil {
		panic(err)
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		panic(errors.New("invalid token"))
	}
	fmt.Println("valid token")
}
func ParseJWTStringWithClaims() {
	jwtKey := make([]byte, 32) // 生成32字节（256位）的密钥

	if _, err := rand.Read(jwtKey); err != nil { //rand.Read函数可以用来生成随机字节序列
		panic(err)
	}
	jwtString := GenerateJWTWithKey(jwtKey)
	claims := jwt.MapClaims{}
	claims["name"] = "hydra2"
	token, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }, jwt.WithExpirationRequired()) //第二个参数是一个回调函数，返回用于验证 JWT 签名的密钥。该函数签名为 func(*Token) (interface{}, error)
	if err != nil {
		panic(err)
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		panic(errors.New("invalid token"))
	}
	fmt.Println("valid token")
	fmt.Println(token.Claims)
}
