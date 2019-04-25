package test

import (
	"math/rand"
	"time"
	"net/http"
	"log"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"encoding/hex"
	sha2562 "crypto/sha256"
	md52 "crypto/md5"
	"os"
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"sort"
	"strconv"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

const year = 2018

func test1()  {
	// result := cal.Add(1,2)
	// fmt.Println(result)

	// result2 := cal.Subtract(110,2)
	// fmt.Println(result2)

	// result3 := multi.Multi(2,4)
	// fmt.Println(result3)

	// var a int = 10
	// fmt.Printf("地址%x\n", &a)
	// a= 11
	// fmt.Printf("地址%x1\n", &a)


	// 创建一个新的结构体
	// book := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	// fmt.Println(book.title)

	//nums := []int{2,3,4}
	//sum := 0
	//for i,num := range(nums) {
	//	sum += num
	//	fmt.Println(i,"==", num)
	//}
	//fmt.Println(sum)
	//
	//kvs := map[string]string{"a": "apple", "b": "banana"}
	//for k,v := range(kvs) {
	//	fmt.Println(k," ", v)
	//}
	//
	//for i, c := range "go" {
	//	fmt.Println(i, c)
	//}
	//
	//fmt.Println(year)
}

/**
  数组操作
 */
func testConst()  {
	// 常量数组长度
	const N = 10
	// 声明数组
	var arr = [N]int{}
	// 随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 循环赋值
	for i:= 0; i< len(arr); i++ {
		// 赋值
		arr[i] = r.Intn(0xff)
	}
	// 打印数组
	fmt.Println(arr)
}

/**
  冒泡排序
 */
func maoPao()  {
	var arr = [11]int{6, 152, 254, 247, 205, 186, 127, 104, 96, 68, 11}
	num := len(arr)
	for i := 0; i< num - 1; i++ {
		for j := 0; j < num - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// 打印数组
	fmt.Println(arr)
}

/**
  md5加密
 */
func md5()  {
	md5Ctx := md52.New()
	md5Ctx.Write([]byte("ms~go"))
	cipherStr := md5Ctx.Sum(nil)
	//cipherStr := md5Ctx.Sum([]byte("h"))
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
}

/**
  sha256
 */
func sha256()  {
	h := sha2562.New()
	h.Write([]byte("ms~go"))
	fmt.Printf("%x", h.Sum(nil))
}


/**
	iota
 */
func iotaTest()  {
	//为 0
	const a = iota
	const (
		//没增加一行+1
		b = iota   // 0
		c = iota   // 1
		d = iota   // 2

		e, f = iota, iota
		// '同一行数值相等'
		g, h, i, j, k = iota, iota, iota, iota, iota
	)
	fmt.Printf("a = %d , b = %d , c = %d , d = %d , e = %d , f = %d , g = %d , h = %d , i = %d , j = %d , k = %d ", a, b, c, d, e, f, g, h, i, j, k)
	//每次 const 出现时，都会让 iota 初始化为0.
	const l = iota

	fmt.Println(l)
}

func iota2()  {
	const(
		a=1<<iota   // 1
		b           // 2
		c			// 4
		d
		e
		f
		_
		g
		h
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		// 异常处理
		panic(e)
	}
}

/**
  文件操作
 */
func fileTest1()  {
	writeString := "测试"
	fileName := "./output1.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(fileName) { // 文件存在
		f, err1 = os.OpenFile(fileName, os.O_APPEND, 0666) // 打开文件
		fmt.Println("文件存在")
	}else {
		f, err1 = os.Create(fileName) //  创建文件
		fmt.Println("文件不存在，开始创建")
	}
	// 检查异常
	check(err1)
	n, err1 := io.WriteString(f, writeString) // 写入文件
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)
}


/**
  文件操作2
 */
func fileTest2()  {
	writeString := "测试"
	fileName := "./output2.txt"
	var d1 = []byte(writeString)
	err2 := ioutil.WriteFile(fileName, d1, 0666)
	check(err2)
}


/**
  文件操作3
 */
func fileTest3()  {
	writeString := "测试"
	fileName := "./output3.txt"
	var d1 = []byte(writeString)
	f, err3 := os.Create(fileName)
	check(err3)
	defer f.Close()
	n2, err3 := f.Write(d1)
	check(err3)
	fmt.Printf("写入 %d 个字节n", n2)
	n3, err3 := f.WriteString("writesn")
	fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()
}

/**
  文件操作4
 */
func fileTest4()  {
	writeString := "测试"
	fileName := "./output3.txt"
	f, err4 := os.Create(fileName)
	check(err4)
	n4, err4 := f.WriteString(writeString)
	fmt.Printf("写入 %d 个字节n", n4)
	check(err4)
	w := bufio.NewWriter(f)
	n4, err4 = w.WriteString("bufferedn")
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}

func mapTest()  {
	//m := make(map[string]string)
	//m["a"] = "a"
	//m["b"] = "b"
	//// 打印字典
	//fmt.Println(m)

	m := map[string]string{}
	m["a"] = "a"
	m["b"] = "b"
	fmt.Println(m)
}

func stringTest()  {
	var print = fmt.Println
	print(strings.Contains("ms~Go", "Ms"))
	print(strings.Join([]string{"a", "b"}, "!"))
	print(strings.Count("ms-go-go-go", "go"))
	print(strings.EqualFold("s", "ss"))
	print(strings.Fields("abc"))
	print(strings.Compare("aa", "ab"))  // aa 比 ab 小 -1， 相等 0 ， 大于 1
}

type ByLength []string

// 我们实现了sort接口的Len，Less和Swap方法
func (s ByLength) Len() int  {
	return len(s)
}

func (s ByLength) Swap(i,j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}

func customerSort()  {
	fruits := []string{"a", "bb", "ccc","dddd","ee","121212"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

func randTest()  {
	var i = 0
	for i < 100 {
		fmt.Println(rand.Intn(0xff))
		i++
	}
}

func exitTest()  {
	var i = 0

	for i < 100 {
		fmt.Println(i)
		if i == 50 {
			//os.Exit(0)
			break
		}

		i++
	}
}

func toUpper()  {
	println(strings.ToUpper("c"))
}

func httpTest()  {
	resp, _ := http.Get("http://www.baidu.com")
	// defer 打开
	defer resp.Body.Close()
	// 获取请求体
	body, _ := ioutil.ReadAll(resp.Body)
	// 打印请求到到数据
	fmt.Println(string(body))
}

func webTest()  {
	http.HandleFunc("/", handler)
	listener := http.ListenAndServe("localhost:9512", nil)
	// 打印输出内容
	// 退出应用程序
	// defer函数不会执行
	log.Fatal(listener)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello world!")
}

/**
  操作excel
 */
func excelTest()  {
	// 创建XMLS文件
	xlsx := excelize.NewFile()
	//循环写入数据
	for i := 1; i < 1000; i++ {
		//拼接字符串
		//位子
		axis := "B" + string(i)
		//在指定的位置写入数据
		xlsx.SetCellValue("Sheet1", axis, 1024*i)

	}
	//创建sheet
	index := xlsx.NewSheet("Sheet2")
	//循环写入数据
	for i := 1; i < 1000; i++ {
		axis := "A" + strconv.Itoa(i)
		//写入
		xlsx.SetCellValue("Sheet2", axis, 1024*i)

	}
	//保存
	xlsx.SetActiveSheet(index)
	//保存文件
	err := xlsx.SaveAs("hello.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
