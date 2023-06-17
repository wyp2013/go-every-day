package pool

import (
	"sync"
	"time"
)

type ConnectPool  interface{
	GetConnect()
	PutConnect()
	Close()
}

type Connect interface{
	IsConnected()
	Close()
}

type connect struct {
	connect     Connect
	updateTime  int
	isIdle      bool
}

type pool struct {
	m           sync.Mutex
	address     []string
	pool        []connect
	timeout     int
	idleTime    int
}

func (p *pool) GetConnect() {
}

func (p *pool) checkIdleConnect() {
	ticker := time.NewTicker(time.Duration(100) * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			p.dealWithIdleConnect()
		}
	}
}

func (p *pool) dealWithIdleConnect() {
	p.m.Lock()
	defer p.m.Unlock()

	// todo 删除空闲的连接
}

func (p *pool) createConnect() {
	// todo create a connect
}

func (p *pool) PutConnect() {
	// todo 收回connect
}

func (p *pool) Close() {
	// todo 关闭所有连接
}

func (p *pool) checkConnection() {
	// todo 检测所有的连接是不是正常连接
}

func (p *pool) IsConnect() bool {
	// todo 检查连接是否正常
	return true
}

