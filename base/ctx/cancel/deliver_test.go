package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)


func goAA(ctx context.Context) {
	newctx, _:= context.WithCancel(ctx)
	ticker := time.NewTicker(1*time.Second)

	for {
		select {
		case <-newctx.Done():
			fmt.Println("goAA exit, msg: ", ctx.Err().Error())
			return

		case <-ticker.C:
			fmt.Println("goAA heart is beating")
		}
	}
}

func goA(ctx context.Context) (context.Context, context.CancelFunc) {
	aCtx, aCancel := context.WithCancel(ctx)

	go goAA(aCtx)

	return aCtx, aCancel
}


func goBB(ctx context.Context) {
	newctx, _:= context.WithCancel(ctx)
	ticker := time.NewTicker(2*time.Second)

	for {
		select {
		case <-newctx.Done():
			fmt.Println("goBB exit, msg: ", ctx.Err().Error())
			return

		case <-ticker.C:
			fmt.Println("goBB heart is beating")
		}
	}
}

func goB(ctx context.Context) (context.Context, context.CancelFunc) {
	bCtx, bCancel := context.WithCancel(ctx)

	go goBB(bCtx)

	return bCtx, bCancel
}


func TestDeliverCancel(t *testing.T) {
	pCtx, _ := context.WithCancel(context.Background())

	_, aCancle := goA(pCtx)
	bCtx, bCancle := goB(pCtx)

	timeA := time.NewTimer(time.Duration(5)*time.Second)
	timeB := time.NewTimer(time.Duration(10)*time.Second)

	for {
		select {
		case <-timeA.C:
			fmt.Println("call a cancel")
			aCancle()
		case <-timeB.C:
			fmt.Println("call B cancel")
			bCancle()
		case <-bCtx.Done():
			return
		}
	}

}

