package api

import (
	"github.com/gin-gonic/gin"
	"gotest/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)                // 注册
	r.POST("/login", login)                      // 登陆
	r.POST("/create_question", create_question)  //问题
	r.POST("/delete_question", delete_question)  //删除问题
	r.POST("/create_answer", create_answer)      //回答
	r.POST("/delete_answer", delete_answer)      //删除回答
	r.POST("/:userid", show_question)            //用户的全部问题
	r.POST("/:userid/:question_id", show_answer) //用户的某问题的回答
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	err := r.Run(":8088")
	if err != nil {
		return
	} // 跑在 8088 端口上
}
