package main

import (
	"fmt"
	"go_net_disk/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	fmt.Println("启动成功")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server,err:%s", err.Error())
	}
}
