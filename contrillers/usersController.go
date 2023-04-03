package contrillers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户列表
var users = []models.User{
	{
		Username: "user1",
		Password: "password1",
	},
	{
		Username: "user2",
		Password: "password2",
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
