/*
@Time : 2019/8/23 14:49
@Author : louis
@File : mian.v0
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/Unknwon/com"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

const tplName  = "http/serve/form.gtpl"

func sayHelloName(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//GET提交的数据会放在URL之后，以?分割URL和传输数据，参数之间以&相连
	fmt.Println(req.Form) //输出到控制台
	fmt.Println("path:", req.URL.Path)
	fmt.Println("scheme:", req.URL.Scheme)
	fmt.Println("string:", req.URL.String())
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello World")
}

func Form(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		if !com.IsExist(tplName) {
			os.MkdirAll(path.Dir(tplName), os.ModePerm)
			os.Create(tplName)
		}
		t, err := template.ParseFiles(tplName)
		if err != nil {
			panic(err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		for k,v :=range r.Form {
			fmt.Println(k,v)
		}
		//必选字段
		if len(r.Form["username"][0])==0 {
			fmt.Fprintf(w,"用户名不能为空")
			return
		}

		//年龄
		if !isInt(r.Form.Get("age")) {
			fmt.Fprintf(w,"年龄必须为正数")
			return
		}

		//中文
		if !isHan(r.Form.Get("hanname")) {
			fmt.Fprintf(w,"请输入中文名")
			return
		}

		//中文
		if !isEng(r.Form.Get("engname")) {
			fmt.Fprintf(w,"请输入正确的英文名")
			return
		}
		//邮箱
		if !isEmail(r.Form.Get("email")) {
			fmt.Fprintf(w,"请输入正确的邮箱")
			return
		}
		//手机号
		if !isMobile(r.Form.Get("mobile")) {
			fmt.Fprintf(w,"请输入正确的手机号")
			return
		}
		//下拉框
		if !isInSingleSelect(r.Form.Get("fruit")) {
			fmt.Fprintf(w,"select not exist")
			return
		}
		//性别
		if !isGender(r.Form.Get("gender")) {
			fmt.Fprintf(w,"性别只能为男女")
			return
		}

		//复选框
		if !isInMultiSelect(r.Form["interest"]) {
			fmt.Fprintf(w,"请选择正确爱好")
			return
		}
		//身份证
		if !isID(r.Form.Get("usercard")) {
			fmt.Fprintf(w,"身份号有误")
			return
		}

		fmt.Fprintf(w,"表单验证成功")
	}
}

func isInt(s string) bool  {
	m,_:= regexp.MatchString(`^[0-9]+$`,s)
	return m
}

func isHan(s string) bool {
	m, _ := regexp.MatchString("^\\p{Han}+$", s)
	return m
}

func isEng(s string) bool {
	m, _ := regexp.MatchString("^[a-zA-Z]+$", s)
	return m
}

func isEmail(s string) bool {
	//`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`
	m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, s)
	return m
}

func isMobile(s string) bool {
	m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, s)
	return m
}

func isInSingleSelect(s string) bool {
	slice := []string{"apple", "pear", "banana"}
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
func isGender(s string) bool {
	slice := []string{"1", "2"}
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func InSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func SliceDiff(slice1, slice2 []string) (diffSlice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

func isInMultiSelect(s []string) bool {
	slice := []string{"football", "basketball", "tennis"}
	if a := SliceDiff(s, slice); a == nil {
		return true
	}
	return false
}

func isID(s string) bool {
	m1, _ := regexp.MatchString(`^(\d{15})$`, s)
	m2, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, s)
	return m1 || m2
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/form", Form)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
访问 http://localhost:9090/?url_long=111&url_long=222  ==> Hello World
控制台输出:
map[url_long:[111 222]]
path: /
scheme:
string: /?url_long=111&url_long=222
[111 222]
key: url_long
value: 111222
*/

// go中的http函数是如何工作的
/*
初始化一个server对象，然后调用了`net.Listen("tcp", addr)`，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。
监控之后如何接收客户端的请求呢？上面代码执行监控端口之后，调用了`srv.Serve(net.Listener)`函数，
这个函数就是处理接收客户端的请求信息。这个函数里面起了一个`for{}`，首先通过`Listener`接收请求，
其次创建一个`Conn`，最后单独开了一个`goroutine`，把这个请求的数据当做参数扔给这个`conn`去服务：`go c.serve()`。
这个就是高并发体现了，用户的每一次请求都是在一个新的`goroutine`去服务，相互不影响。

那么如何具体分配到相应的函数来处理请求呢？`conn`首先会解析request:`c.readRequest()`,
然后获取相应的handler:`handler := c.server.Handler`，也就是我们刚才在调用函数`ListenAndServe`时候的第二个参数，
我们前面例子传递的是nil，也就是为空，那么默认获取`handler = DefaultServeMux`,那么这个变量用来做什么的呢？
对，这个变量就是一个路由器，它用来匹配url跳转到其相应的handle函数，那么这个我们有设置过吗?
有，我们调用的代码里面第一句不是调用了`http.HandleFunc("/", sayhelloName)`嘛。
这个作用就是注册了请求`/`的路由规则，当请求uri为"/"，路由就会转到函数`sayhelloName`，`DefaultServeMux`会调用`ServeHTTP`方法，
这个方法内部其实就是调用`sayhelloName`本身，最后通过写入`response`的信息反馈到客户端。
*/
