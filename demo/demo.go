package demo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = []byte("mysecret") // JWT密钥
)

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

// JWT认证结构体
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


// 用户列表
var users = []User{
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
func registerHandler(c *gin.Context) {
	var registerRequest RegisterRequest
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

	// 创建新用户
	newUser := User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}
	users = append(users, newUser)

	// 生成JWT Token
	expirationTime := time.Now().Add(24 * time.Hour) // Token有效期为24小时
	claims := &Claims{
		Username: registerRequest.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


// 登录接口
func loginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户名和密码
	for _, user := range users {
		if user.Username == loginRequest.Username && user.Password == loginRequest.Password {
			// 生成JWT Token
			expirationTime := time.Now().Add(24 * time.Hour) // Token有效期为24小时
			claims := &Claims{
				Username: loginRequest.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// 需要认证的路由组
func authenticatedRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取JWT Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		// 解析JWT Token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// 验证JWT Token
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// 将用户名保存到上下文中
		c.Set("username", claims.Username)
		c.Next()
	}
}

// 检查当前登录用户的路由
func currentUserHandler(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}





// 注册接口
func registerHandler(c *gin.Context) {
	var registerRequest RegisterRequest
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

	// 创建新用户
	newUser := User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}
	users = append(users, newUser)

	// 生成JWT Token
	expirationTime := time.Now().Add(24 * time.Hour) // Token有效期为24小时
	claims := &Claims{
		Username: registerRequest.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 注册路由
r.POST("/register", registerHandler)












func demo() {
	// 创建Gin Engine
	r := gin.Default()

	// 注册登录接口
	r.POST("/login", loginHandler)

	// 需要认证的路由组
	authGroup := r.Group("/")
	authGroup.Use(authenticatedRoutes())
	{
		// 检查当前登录用户的路由
		authGroup.GET("/currentUser", currentUserHandler)
	}

	// 启动HTTP服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
