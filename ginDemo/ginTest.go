package ginDemo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RequestTest() {

	router := gin.Default()
	//router.Use(Logger())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
		//c.String(http.StatusOK, "this is 200")
	})

	//解析路径参数
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})
	//解析Query参数
	router.GET("users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	//POST
	router.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		//password := c.PostForm("password")
		password := c.DefaultPostForm("password", "00000000")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//Query 和 POST混合
	router.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultQuery("password", "00000000")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	//Map参数
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	//json
	router.POST("/json", func(c *gin.Context) {
		json := User{}
		//转换请求体
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Msg": err.Error(),
			})
			return
		}

		log.Println("request Param ", json)

		c.JSON(http.StatusOK, gin.H{
			"username": json.Username,
			"password": json.Password,
		})
	})

	//Redirect
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/user/abc")
	})

	router.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		router.HandleContext(c)
	})

	//group routes
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	//v1
	v1 := router.Group("/v1")
	{
		v1.GET("posts", defaultHandler)
		v1.GET("series", defaultHandler)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("posts", defaultHandler)
		v2.GET("series", defaultHandler)
	}

	//router.Use(gin.Recovery())

	router.Run(":1080")
}

//func Logger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		c.Set("hello", "world")
//		c.Next()
//		latency := time.Since(t)
//
//		url := c.FullPath()
//
//		remoteIp := c.RemoteIP()
//		localIP := c.ClientIP()
//
//		log.Println(latency, "----", url, "ip: ", localIP+"----->  "+remoteIp)
//	}
//}
