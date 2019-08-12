package orchan

import (
	"fmt"
	"sync"
)

func OrChan(chans ...chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func () {
		var once sync.Once
		for _, ch := range chans {
			go func(c chan interface{}) {
				select {
				case x, ok := <-ch:
					if ok {
						fmt.Println(x)
					}

					once.Do(func() {close(c)})
				case <-out:
				}
			}(out)
		}
	}()

	return out
}

func or(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	}

	done := make(chan interface{})
	go func() {
		defer close(done)
		switch len(chans) {
		case 2:
			select {
			case <-chans[0]:
			case <-chans[1]:
			}

		default:
			m := len(chans)/2
			select {
			case <-or(chans[:m]...):
			case <-or(chans[m:]...):
			}
		}
	}()

	return done
}
