package api

import (
	"fmt"
	"time"
)

func Loading(done chan bool) {
	for {
		select {
		case _, ok := <-done:
			if !ok {
				fmt.Println("mysql数据库已连接，检查表结构中...")
				return
			}
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
