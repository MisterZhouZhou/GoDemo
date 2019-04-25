package httpT

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"../define"
)

//用于保存用户账号信息（注意，结构体里面的属性必须首字母大写）
type User struct {
	Id       uint        `json:"id"`
	Username string 	 `json:"username"`
	Password string 	 `json:"-"`          //其中`json:"-"`代表忽视此属性
	CreateTime time.Time `json:"create_time"`
}

var userArr = make([]User, 0) // 用于存储用户的信息

//查看用户名是否已经存在
func Existed(user User) bool {
	for _, value := range userArr {
		if value.Username == user.Username { //判断用户名是否存在
			return true
		}
	}
	return false
}

//验证用户输入的是否正确
func Verify(user User) bool {
	for _, value := range userArr {
		// 判断用户名与密码是否都相同
		if value.Username == user.Username && value.Password == user.Password {
			return true
		}
	}
	return false
}

// 注册, 参数必须是以json形式请求
func Register(userInfo []byte) define.CommonResult {
	var user User
	json.Unmarshal(userInfo, &user)//将json转换成结构体
	if Existed(user) { //判断是否已经注册过
		return define.CustomerCommonResult(-1, "用户名已存在", nil)
	}
	if len(userArr) >= 1 {
		lastUser := userArr[len(userArr)-1]
		user.Id = lastUser.Id + 1
	}else {
		user.Id = 1
	}
	user.CreateTime = time.Now()
	userArr = append(userArr, user) //将这个用户的信息保存到切片中
	//return define.CommonResult{
	//	BaseResult: define.Success,
	//	Data:       "注册成功",
	//}
	return define.CustomerCommonResult(0, "注册成功", user)
}


// 登录
func LoginIn(userInfo []byte) define.CommonResult {
	var user User
	json.Unmarshal(userInfo, &user)
	if !Existed(user) { // //首先判断用户输入的用户名是否存在
		return define.CustomerCommonResult(-1, "用户名不存在", nil)
	}
	if !Verify(user) {
		return define.CustomerCommonResult(-1, "用户名或密码错误",nil)
	}
	return define.CustomerCommonResult(0, "登录成功", user)
}




//登录的具体交互函数
func login(res http.ResponseWriter, req *http.Request)  {
	if req.Method == "POST" {
		s, _ := ioutil.ReadAll(req.Body)
		// 登录
		res.Write(FeeBookCommonResultToByte(LoginIn(s)))
	}else {
		result := define.CustomerCommonResult(-1, "只支持POST方式",nil)
		res.Write(FeeBookCommonResultToByte(result))
	}
}


//注册的具体交互函数
func register(res http.ResponseWriter, req *http.Request)  {
	if req.Method == "POST" { //判断是不是POST请求
		s, _ := ioutil.ReadAll(req.Body) // 读取数据
		res.Write(FeeBookCommonResultToByte(Register(s)))
	}else {
		result := define.CustomerCommonResult(-1, "只支持POST方式",nil)
		res.Write(FeeBookCommonResultToByte(result))
	}
}

func RunServer()  {
	http.HandleFunc("/login", login)  		// 登录
	http.HandleFunc("/register", register)  // 注册
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("服务启动失败")
	}
}
