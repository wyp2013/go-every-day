package transmit

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

func transContext() {
	ctx := context.TODO()
	req := map[string]string{}
	req["traceId"] = string(uuid.NewUUID())

	getValueWithContext(ctx, req)
}


func getValueWithContext(ctx context.Context, req map[string]string) {
	traceId := ""
	if val, ok := req["traceId"]; ok {
		traceId = val
	}

	ctx = context.WithValue(ctx, "traceId", traceId)

	go func() {
		traceId, ok:= ctx.Value("tranceId").(string)
		if ok {
			fmt.Println(traceId)
		} else {
			fmt.Println("can not find key traceId")
		}

		time.Sleep(time.Second * 1)
	}()
}


func calculatePos() {}
func sendPos() {}

func doSendPos(ctx context.Context) {
	timer := time.NewTicker(time.Second)
	for {
		calculatePos()
		sendPos()

		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			// 阻塞1s
		}
	}
}

func userExit() bool {
	return false
}

func cancelFunc() {
	ctx, cancle := context.WithCancel(context.TODO())
	go doSendPos(ctx)

	// do other thing

	// 获取用户是否取消
	exit := userExit()
	if exit {
		cancle()
	}
}

func CreateCancel() map[string]context.CancelFunc {
	tree := make(map[string]context.CancelFunc)
	mctx, mcan := context.WithCancel(context.TODO())
	tree["mctx"] = mcan

	sy := sync.Mutex{}


	go func() {
		tickerM := time.NewTicker(time.Second * 5)


		// a
		go func() {
			actx, acan := context.WithCancel(mctx)
			sy.Lock()
			tree["actx"] = acan
			sy.Unlock()

			tickerA := time.NewTicker(time.Second * 5)

			//c
			go func() {
				cctx, ccan := context.WithCancel(actx)
				sy.Lock()
				tree["cctx"] = ccan
				sy.Unlock()

				tickerC := time.NewTicker(time.Second * 2)
				for {
					select {
					case <-cctx.Done():
						fmt.Println("goroutine c cancel is called")
						return
					case <-tickerC.C:
						fmt.Println("c goroutine is beating")
					}
				}

			}()

			for {
				select {
				case <-actx.Done():
					fmt.Println("a cancel is called, a is exit")
					return
				case <-tickerA.C:
					fmt.Println("goroutine a is beating")
				}
			}

		}()



		// b
		go func() {
			tickerB := time.NewTicker(time.Second * 3)
			bctx, bcan := context.WithCancel(mctx)

			sy.Lock()
			tree["bctx"] = bcan
			sy.Unlock()



			// d
			go func() {
				tickerD := time.NewTicker(time.Second * 1)
				dctx, dcan := context.WithCancel(bctx)

				sy.Lock()
				tree["dctx"] = dcan
				sy.Unlock()

				for {
					select {
					case <-dctx.Done():
						fmt.Println("goroutine d cancel is called and exit")
					case <-tickerD.C:
						fmt.Println("goroutine d is running and is beating")
					}
				}
			}()


			// e
			go func() {
				tickerE := time.NewTicker(time.Second * 2)
				ectx, ecan := context.WithCancel(bctx)

				sy.Lock()
				tree["ectx"] = ecan
				sy.Unlock()

				for {
					select {
					case <-tickerE.C:
						fmt.Println("goroutine e is running and is beating")
					case <-ectx.Done():
						fmt.Println("goroutine e cancel is called and exit")
					}
				}
			}()




			for {
				select {
				case <-bctx.Done():
					fmt.Println("bcan is called and goroutine b is exist")
					return

				case <-tickerB.C:
					fmt.Println("goroutine is running and is beating")
				}
			}
		}()


		for {
			select {
			case <-mctx.Done():
				fmt.Println("mcan is called, and createCancel is Exit")
				return
			case <-tickerM.C:
				fmt.Println("goroutine m is beating")
				fmt.Println("is liveing")
			}
		}
	}()

	return tree
}


func getContext() {

}



