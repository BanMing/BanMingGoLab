package main

import (
	"html/template"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template") //创建一个模板
	//解析一个模板文件
	t, _ = t.ParseFiles("tmpl/welcome.html")
	//获取当前用户信息
	//user:=GetUser()
	//执行模板的merger操作
	t.Execute(w, net.Interface{})
}
func main() {
	//test1()
	//test2()
	test3()
}

type PersonT struct {
	UserName string
}
type PersonF struct {
	UserName string
	Emails   []string
	Friend   []*Friend
}
type Friend struct {
	Fname string
}

func test1() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := PersonT{UserName: "BanMing"}
	t.Execute(os.Stdout, p)
}

func test2() {
	f1 := Friend{Fname: "ssss"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname.example")
	t, _ = t.Parse(`hello {{.UserName}}!
			 {{range .Emails}}
			 an email {{.}}
			 {{end}}
			 {{with .Friends}}
			 {{range .}}
			 my friend name is {{.Fname}}
			 {{end}}
			 {{end}}`)
	p := PersonF{UserName: "BanMing", Emails: []string{"asdasda.com", "asdasdas.com"}, Friend: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}

func test3() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo:{{if ``}} 不会输出 .{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)
	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipline if demo {{if `anything`}} 我有内容,我会输出 .{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if 部分{{else}} else部分 .{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

}
