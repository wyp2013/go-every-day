## channel

### send channel
ch <- x

发送channel
```
1. 如果channel的recvq不为空，则直接把message拷贝到recvq队列首部的sg，并且把sg从队列上删除，
把sg设置为可运行状态
2. 如果channel带缓冲，并且缓冲区没有满，直接把message放到缓冲区
3. 如果缓冲区满，则阻塞当前gorouting，构建一个sd放在sendq队列
```
### recv channel
y <- ch

接收channel
```
1. 如果 Channel 是空的，那么就会直接调用 gopark 挂起当前的 Goroutine；
2. 如果 Channel 已经关闭并且缓冲区没有任何数据，chanrecv 函数就会直接返回；
3. 如果 Channel 上的 sendq 队列中存在挂起的 Goroutine，就会将recvx 索引所在的数据拷贝到接收变量所在的内存空间上并将 sendq 队列中 Goroutine 的数据拷贝到缓冲区中；
4. 如果 Channel 的缓冲区中包含数据就会直接从 recvx 所在的索引上进行读取；
5. 在默认情况下会直接挂起当前的 Goroutine，将 sudog 结构加入 recvq 队列并更新 Goroutine 的 waiting 属性，最后陷入休眠等待调度器的唤醒；
```