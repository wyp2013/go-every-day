package cancel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)


func createTest() map[string]context.CancelFunc {
	tree := make(map[string]context.CancelFunc)
	mctx, mcan := context.WithCancel(context.TODO())
	tree["mctx"] = mcan

	sy := &sync.Mutex{}

	go func() {
		ticker := time.NewTicker(5*time.Second)

		// a
		go func() {
			actx, acan := context.WithCancel(mctx)
			sy.Lock()
			tree["actx"] = acan
			sy.Unlock()
			tickerA := time.NewTicker(4*time.Second)

			//c
			go func() {
				cctx, ccan := context.WithCancel(actx)
				sy.Lock()
				tree["cctx"] = ccan
				sy.Unlock()
				tickerC := time.NewTicker(2*time.Second)

				for {
					select {
					case <-cctx.Done():
						fmt.Println("ccan called")
						return
					case <-tickerC.C:
						fmt.Println("c goroutine heart is beating")
					}
				}
			}()

			for {
				select {
				case <-actx.Done():
					fmt.Println("acan called")
					return
				case <-tickerA.C:
					fmt.Println("a goroutine heart is beating")
				}
			}
		}()


		// b
		go func() {
			bctx, bcan := context.WithCancel(mctx)
			sy.Lock()
			tree["bctx"] = bcan
			sy.Unlock()
			tickerB := time.NewTicker(3*time.Second)

			// d
			go func() {
				dctx, dcan := context.WithCancel(bctx)
				sy.Lock()
				tree["dctx"] = dcan
				sy.Unlock()
				tickerD := time.NewTicker(1*time.Second)

				for {
					select {
					case <-dctx.Done():
						fmt.Println("dcan called")
						return
					case <-tickerD.C:
						fmt.Println("d goroutine heart is beating")
					}
				}
			}()

			for {
				select {
				case <-bctx.Done():
					fmt.Println("bcan called")
					return
				case <-tickerB.C:
					fmt.Println("b goroutine heart is beating")
				}
			}
		}()

		for {
			select {
			case <-mctx.Done():
				fmt.Println("mcan called")
				return
			case <-ticker.C:
				fmt.Println("m goroutine heart is beating")
			}
		}
	}()

	return tree
}

// 父节点cancel，子节点都cancel
func TestRootCancel(t *testing.T) {
	tree := createTest()
	time.Sleep(5*time.Second)
	tree["mctx"]()

	// 程序打印完
	time.Sleep(6*time.Second)
}

// 子节点都cancel
func TestCancel2(t *testing.T) {
	tree := createTest()
	time.Sleep(5*time.Second)
	tree["actx"]()

	// 让b分支运行
	time.Sleep(20*time.Second)
}

// 子节点都cancel
func TestCancel3(t *testing.T) {
	tree := createTest()
	time.Sleep(5*time.Second)
	tree["dctx"]()

	// 除了d，其它都运行
	time.Sleep(20*time.Second)
}
