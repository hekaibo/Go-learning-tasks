package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	defer fmt.Println("退出")
	for {

		taskID := rand.Intn(100)
		var msg string
		var taskTime int
		fmt.Println("请输入消息内容：")
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		fmt.Println("请输入延迟时间：")
		fmt.Scanln(&taskTime)

		url := fmt.Sprintf("http://127.0.0.1:8888/add-task?ID=%d&Message=%s&Time=%d", taskID, msg, taskTime)
		//fmt.Println(url)
		fmt.Printf("任务 %d 已发送\n", taskID)

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		fmt.Println("************************************")
	}
}
