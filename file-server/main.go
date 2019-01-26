package main

import "net/http"

func main() {
	//web监听8080端口
	//设置当前目录为文件系统根目录
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
