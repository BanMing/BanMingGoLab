package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//xml 文件大结构
type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	//innerxml xml中对应的所有数据
	Description string `xml:",innerxml"`
}

//服务器结构
type server struct {
	//XMLName不会被输出
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	WriteXmlFile()
}

func WriteXmlFile() {
	v := &Recurlyservers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	//这里生成xml字符不包括头
	output, err := xml.MarshalIndent(v, "	", "	")
	if err != nil {
		fmt.Println("err:", err)
	}
	//这里插入标准xml头文件
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
func ReadXmlFile() {
	file, err := os.Open("/Users/BanMing/Coding/Go/BanMingGoLab/TextLearn/XML/servers.xml")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	v := Recurlyservers{}
	//把指定二进制数据转化成指定数据结构
	//使用结构体的tag来反射转至
	//如果没有找到tag再由字段名来找
	//大小写敏感
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	//fmt.Println(v)
	fmt.Println(v.XMLName)
}
