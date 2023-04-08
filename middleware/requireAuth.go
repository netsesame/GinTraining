package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ParseToken 解析JWT
// 该函数接受两个参数，第一个是JWT字符串，第二个是用于验证签名的密钥。它使用JWT库的Parse函数来解析JWT，
// 如果签名验证失败，则返回错误。如果签名验证通过，则返回JWT的声明，类型为jwt.MapClaims。
// 在Parse函数中，我们使用了一个回调函数来验证签名。该函数返回一个密钥用于验证签名。
// 在这个例子中，我们使用HMAC算法来签名，所以我们需要提供一个密钥来验证签名。如果签名算法不是HMAC，我们会返回一个错误。
// 最后，我们检查JWT是否有效，并返回声明。如果JWT无效，则返回错误。
func ParseToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWtSECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	// JWT无效
	return nil, fmt.Errorf("invalid token")
}

// 中间件
func ReqireAuth(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	// 解码/验证它
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	claims, err := ParseToken(tokenString, []byte(os.Getenv("JWtSECRET")))
	if err != nil {
		// JWT无效
		fmt.Println("Error parsing token: ", err)
		return
	}

	fmt.Println("Token claims: ", claims, claims["sub"])
	// 获取过期时间 /检查exp
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		//JWT已过期
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "JWT已过期"})
	} else {
		// JWT未过期
		fmt.Println("还没过期")
	}
	data := map[string]interface{}{
		"token": tokenString,
		"user":  claims["user"],
		"exp":   claims["exp"],
	}
	// 找到带有token sub的用户
	c.Set("data", data)

	// 附加到req

	// 继续
	c.Next()

}
