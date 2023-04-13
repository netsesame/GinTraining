package contrillers

import (
	"fmt"
	"main/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户列表
var users = []models.User{
	{
		Username: "11",
		Password: "$2a$10$DpNjNVsDnPia1xp/Zq09Me9yf4lipAdm9zyT5kHxSU.XmkzWA2zrS",
	},
	{
		Username: "33",
		Password: "$2a$10$mXyrzcQ7EZfoggRLwqkPMu.z5L2BuTe7gXL7fFpxGmbEsrUZ5S/S2",
	},
}

// 注册接口
func RegisterHandler(c *gin.Context) {
	var registerRequest models.RegisterRequest
	//请求体中的JSON对象绑定到指定的结构体中 ,直接JSON 格式的数据
	//c.Bind()方法会根据请求头中的Content-Type自动解析请求体中的数据格式，
	//并将其绑定到指定的结构体中。如果请求体中的数据格式不是JSON格式，
	//c.Bind()方法会尝试根据其他类型进行解析，例如XML或form数据。如果解析失败，将返回错误。
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断用户名是否已经存在
	for _, user := range users {
		if user.Username == registerRequest.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
	}

	// 对密码进行bcrypt加密存储
	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//添加到模拟数据库
	addUserDb := models.User{Username: registerRequest.Username, Password: string(hash)}
	users = append(users, addUserDb)
	fmt.Println(users)
	// // 生成JWT Token
	// expirationTime := time.Now().Add(24 * time.Hour) // Token有效期为24小时
	// claims := &models.Claims{
	// 	Username: registerRequest.Username,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(jwtSecret)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"token": "注册成功"})
}

// 登录接口
func LoginHandler(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户名和密码
	for _, user := range users {
		// && user.Password == loginRequest.Password
		fmt.Println("数据库里的用户名", user.Username, "前端输入用户名", loginRequest.Username)
		if user.Username == loginRequest.Username {
			// 生成JWT Token

			expirationTime := time.Now().Add(24 * time.Hour) // Token有效期为24小时
			// 创建一个我们自己的声明
			claims := &models.Claims{
				Username: loginRequest.Username, // 自定义字段
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			// 使用指定的签名方法创建签名对象
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//tokenString, err := token.SignedString(jwtSecret)
			//在 JWT 中，签名密钥是一个字节数组，而不是字符串。因此，
			//需要将字符串类型的签名密钥转换为字节数组类型。可以使用
			//[]byte() 函数将字符串转换为字节数组。例如，如果签名密钥是一个名为 jwtSecret 的字符串变量，
			//则可以使用 []byte(jwtSecret) 将其转换为字节数组类型。
			tokenString, err := token.SignedString([]byte(os.Getenv("JWtSECRET")))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.SetSameSite(http.SameSiteDefaultMode)

			cookieExpirationTime := time.Now().Add(24 * time.Hour)
			cookie := &http.Cookie{
				Name:     "Authorization",
				Value:    tokenString,
				Expires:  cookieExpirationTime,
				HttpOnly: true,
				SameSite: http.SameSiteDefaultMode,
			}
			c.SetCookie(
				cookie.Name,
				cookie.Value,
				cookie.MaxAge,
				cookie.Path,
				cookie.Domain,
				cookie.Secure,
				cookie.HttpOnly,
			)

			data := models.LoginResponse{
				Token:    tokenString,
				Username: loginRequest.Username,
			}

			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": data})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
// 这里假设Token放在Header的Authorization中，并使用Bearer开头
// 这里的具体实现方式要依据你的实际业务情况决定

func Validate(c *gin.Context) {
	dataTab, _ := c.Get("data")
	c.JSON(http.StatusOK, gin.H{
		"message": dataTab,
	})
}

// 返回支持的全部交易所JSON
func GetExchangesHandler(c *gin.Context) {
	exchanges := map[string]string{
		"Binance": "币安",
		"Huobi":   "火必",
		"Okex":    "欧易",
		"Gate":    "芝麻开门",
	} // 假设支持这些交易所
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Success", "data": exchanges})
}
