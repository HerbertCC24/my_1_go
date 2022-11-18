package main

import "github.com/gin-gonic/gin"
import "encoding/json"
// import "fmt"

type User struct {
	username string
	password  string
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":  200,
			"data": []string{"Herbert","Annie"},
			"message": "pong",
		})
	})
	// 传统形式
	r.GET("/user/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		c.JSON(200, gin.H{
			"username" : username,
			"password" : password,
			"code": 200,
			"data": nil,
			"message": "Succeed!",
		})
	})
	// restful
	r.GET("/user/login/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		c.JSON(200, gin.H{
			"username" : username,
			"password" : password,
			"code": 200,
			"data": nil,
			"message": "Succeed!",
		})
	})
	//前端向后端传递json
	r.POST("/json", func(c *gin.Context) {
		data,_ := c.GetRawData()
		var m map[string]interface{}
		//包装为json数据
		_ = json.Unmarshal(data,&m)
		c.JSON(200, m)
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}