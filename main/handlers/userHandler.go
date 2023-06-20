package handlers

import (
	"LoginAndRegisterDemo/main/models"
	"LoginAndRegisterDemo/main/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct{}

// 注册页面
func (u UserHandler) Register(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", "")
}

// 接收注册表单，处理注册请求
func (u UserHandler) DoRegister(ctx *gin.Context) {
	name := ctx.PostForm("username")
	salt := utils.NewSalt()
	passwd := utils.EncPasswd(ctx.PostForm("passwd"), salt)
	user := &models.User{
		UserName: name,
		Password: passwd,
		Salt:     salt,
	}
	db, exists := ctx.Get("DB")
	if !exists {
		// 新建一个 gorm 实例
		log.Fatal("Get('DB') error")
	}
	_, exists, err := models.GetUser(db.(*gorm.DB), name)
	if err != nil {
		// 查询发生错误
		log.Fatalf("GetUser() error: %s", err.Error())
	}
	if exists {
		// 用户已存在
		log.Fatal("用户已存在")
	}
	if err := models.AddUser(db.(*gorm.DB), user); err != nil {
		// 插入数据错误
		log.Fatalf("AddUser error: %s", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"用户名": name,
		"状态":  "注册成功",
	})
}

// 登录页面
func (u UserHandler) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", "")
}

// 接收登录表单，处理登录请求
func (u UserHandler) DoLogin(ctx *gin.Context) {
	name := ctx.PostForm("username")
	passwd := ctx.PostForm("passwd")
	db, exists := ctx.Get("DB")
	if !exists {
		// 新建一个 gorm 实例
		log.Fatal("Get('DB') error")
	}

	user, exists, err := models.GetUser(db.(*gorm.DB), name)
	if err != nil {
		// 查询发生错误
		log.Fatalf("GetUser() error: %s", err.Error())
	}
	if exists {
		if utils.EncPasswd(passwd, user.Salt) == user.Password {
			// 生成 token String
			tokenStr, err := utils.GenerateToken(user.UserName, user.Password)
			if err != nil {
				log.Fatalf("生成token字符串发生错误: %s", err.Error())
			}
			ctx.SetCookie("Authorization", tokenStr, 60, "/", "127.0.0.1", false, true)
			// ctx.String(http.StatusOK, "登录成功，%v的token已设置，即将跳转到主页", name)
			// 重定向到主页
			ctx.Redirect(http.StatusMovedPermanently, "/")
		} else {
			ctx.String(http.StatusOK, "密码错误")
		}
	} else {
		// 用户不存在
		ctx.String(http.StatusOK, "用户不存在")
	}
}

// 忘记密码链接，进入修改密码页面
func (u UserHandler) ChangePwd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "changePwd.html", "")
}

// 接收修改密码表单，处理修改密码请求
func (u UserHandler) DoChangePwd(ctx *gin.Context) {
	name := ctx.PostForm("username")
	salt := utils.NewSalt()
	passwd := utils.EncPasswd(ctx.PostForm("newPasswd"), salt)
	//rePwd := utils.EncPasswd(ctx.PostForm("rePasswd"), salt)
	db, exists := ctx.Get("DB")
	if !exists {
		// 新建一个 gorm 实例
		log.Fatal("Get('DB') error")
	}
	user := &models.User{
		UserName: name,
		Password: passwd,
		Salt:     salt,
	}
	if err := models.UpdatePasswd(db.(*gorm.DB), user); err != nil {
		ctx.String(http.StatusOK, "修改失败")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"用户名": name,
			"状态":  "修改成功",
		})
	}
}
