package cancel

import (
	"context"
	"testing"
	"time"
)


func calculatePos() {}
func sendPos() {}
func userExit() bool {return false}

func do(ctx context.Context) {
	for {
		calculatePos()
		sendPos()
		select {
		case <-ctx.Done():
			// 用户退出页面，则退出程序
			return
		case <-time.After(time.Second):
			// 阻塞1s
		}

	}
}

func TestCancelFun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())

	// 发送给用户外卖小哥位置信息
	go do(ctx)

	// do other thing

	// 获取用户是否退出
	exit := userExit()
	if exit {
		cancel()
	}
}
