package models

import (
	"github.com/dgrijalva/jwt-go"
	//"gorm.io/gorm"
)

// // 账号密码的结构体
// type User struct {
// 	gorm.Model
// 	Username string
// 	Password string
// }

// 用户信息结构体
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT认证结构体
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 登录返回的结构体
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
