package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//////////////////////////////       UserController        ///////////////////////////////
// 基本的GET请求获取用户名
type UserController struct {
	beego.Controller
}

type Date struct {
	name string
	age  int
	sex  string
}

func (c *UserController) Get() {
	data := Date{"magic", 24, "男"}
	content := strings.Join([]string{"姓名:" + data.name, "性别:" + data.sex, "年龄:" + strconv.Itoa(data.age)}, " ")
	c.Ctx.WriteString(content)
}

//////////////////////////////       LoginController        ///////////////////////////////
// 基本的Post请求
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	content := strings.Join([]string{"用户名:" + username, "密码:" + password}, " ")
	c.Ctx.WriteString(content)
}

//////////////////////////////       JSONController        ///////////////////////////////
// 基本的Post请求
type JSONController struct {
	beego.Controller
}

type JSONStruct struct {
	Code int
	Msg  string
}

// 属性名必须为大些, 可以设置tag来进行重命名
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func (c *JSONController) Get() {
	// mystruct := &JSONStruct{0, "hello"}
	group := []Person{
		Person{"zw", 24, "man"},
		Person{"zw2", 25, "man"},
	}
	// mystruct := &Person{"zw", 24, "man"}
	c.Data["json"] = &group
	c.ServeJSON()
}
