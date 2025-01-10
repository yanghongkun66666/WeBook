package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

//别人的直接薅过来用就行

type UserHandler struct {
	//userhandler定义所有跟用户有关的路由
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	//初始化的时候处理好
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
		//go 标准正则库不支持这种语法，能力非常有限

		userIdKey = "userId"
		bizLogin  = "login"
	)
	//什么时候用到就定义正则表达式
	//就近原则，最小化作用域原则
	//写回响应
	//预编译
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}

}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	//这是进来就已经分好组了
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//进一步拆分
	//handlers定义在哪里，就拿handlers去注册
	//注册

	//分组路由
	//   /users/profile 不算前缀路由 就是静态路由
	ug := server.Group("/users")
	fmt.Println(u)      // 打印 UserHandler 实例
	fmt.Println(server) // 打印 gin.Engine 实例
	fmt.Println(ug)     // 打印 RouterGroup 实例

	ug.POST("/signup", u.SignUp)
	ug.GET("/profile", u.Profile)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)

	//server.POST("/users/signup", u.SignUp)
	//
	////登录
	//server.POST("/users/login", u.Login)
	//
	////编辑用户
	//server.POST("/users/edit", u.Edit)
	//
	//server.GET("users/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	//方法的接收者是 userHandler 类型的指针，表示该方法是与
	//userHandler 结构体关联的。也就是说，该方法是 userHandler 类型的实例上的一个行为。

	//定义结构体来接收前端传过来的字段
	//内部结构体，我不想别的方法能够看到
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		//这个confirm完全就可以在前端就验证了
	}

	var req SignUpReq
	//Bind 方法会根据Content-Type来解析你的数据到req里面
	//解析没过就会直接写会 400的错误，这些都不需要你去管
	if err := ctx.Bind(&req); err != nil {
		return
		//确保你req拿到的就是前端传过来的数据
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
		//err.Error()不能返回，因为这是你的内部信息,返回一种别人猜不到你的系统发生了什么事情
	}

	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//这里按道理应该记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
		//err.Error()不能返回，因为这是你的内部信息,返回一种别人猜不到你的系统发生了什么事情
	}

	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	ctx.String(http.StatusOK, "注册成功")

	fmt.Printf("%v", req)
	//数据库操作
}

func (u *UserHandler) Login(ctx *gin.Context) {
	//方法的接收者是 userHandler 类型的指针，表示该方法是与
	//userHandler 结构体关联的。也就是说，该方法是 userHandler 类型的实例上的一个行为。

}

func (u *UserHandler) Edit(ctx *gin.Context) {
	//方法的接收者是 userHandler 类型的指针，表示该方法是与
	//userHandler 结构体关联的。也就是说，该方法是 userHandler 类型的实例上的一个行为。

}

func (u *UserHandler) Profile(ctx *gin.Context) {
	//方法的接收者是 userHandler 类型的指针，表示该方法是与
	//userHandler 结构体关联的。也就是说，该方法是 userHandler 类型的实例上的一个行为。

}
