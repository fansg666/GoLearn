package main

import "github.com/gin-gonic/gin"

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello goland!",
	})
}
func main() {
	r := gin.Default() //返回默认的路有引擎
	//指定用户使用get请求访问/hello时,执行sayHello这个函数
	r.GET("/hello", sayhello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")
}
