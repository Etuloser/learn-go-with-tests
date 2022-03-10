// 使用 http 构建静态文件服务器
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func FileServer() {
	// Args保管了命令行参数，第一个是程序名。
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(p)
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":10486", nil)
	if err != nil {
		fmt.Println(err)
	}
}
