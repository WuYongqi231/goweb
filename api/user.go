package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gotest/api/middleware"
	"gotest/dao"
	"gotest/model"
	"gotest/utils"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}
	// 传入用户名和密码

	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否重复
	flag := dao.SelectUser(username)

	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user already exists")
		return
	}

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}

	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func create_question(c *gin.Context) {

	user_id := c.PostForm("user_id")
	content := c.PostForm("content")
	flag := dao.AddQuestion(user_id, content)
	if !flag {
		utils.RespFail(c, "create question failed")
		return
	}
	utils.RespSuccess(c, "create question successful")
}
func delete_question(c *gin.Context) {
	question_id := c.PostForm("question_id")
	flag := dao.DeleteQuestion(question_id)
	if !flag {
		utils.RespFail(c, "delete question failed")
	}
	utils.RespSuccess(c, "delete question successful")
}
func create_answer(c *gin.Context) {
	question_id := c.PostForm("question_id")
	user_id := c.PostForm("user_id")
	content := c.PostForm("content")
	flag := dao.CreateAnswer(question_id, content, user_id)
	if !flag {
		utils.RespFail(c, "create answer failed")
	}
	utils.RespSuccess(c, "create answer successful")
}
func delete_answer(c *gin.Context) {
	answer_id := c.PostForm("answer_id")
	flag := dao.DeleteAnswer(answer_id)
	if !flag {
		utils.RespFail(c, "delete question failed")
	}
	utils.RespSuccess(c, "delete answer successful")
}
func show_question(c *gin.Context) {
	userid := c.Param("userid")
	flag, ans := dao.ShowQuestion(userid)
	if !flag {
		utils.RespFail(c, ans)
	}
	utils.RespSuccess(c, ans)
}

func show_answer(c *gin.Context) {
	asker := c.Param("userid")
	question_id := c.Param("question_id")
	flag, ans := dao.ShowAnswer(asker, question_id)
	if !flag {
		utils.RespFail(c, ans)
	}
	utils.RespSuccess(c, ans)
}
