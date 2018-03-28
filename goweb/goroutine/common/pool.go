package common

import (
	"errors"
	"io"
	"log"
	"sync"
)

var ErrPoolClosed = errors.New("资源池已关闭")

//一个安全的资源池，被管理的资源必须都实现io.Close接口
type Pool struct {
	m       sync.Mutex                //互斥锁变量
	res     chan io.Closer            //一个有缓冲的通道，用来保存共享的资源
	factory func() (io.Closer, error) //一个函数类型，它的作用就是当需要一个新的资源时，可以通过这个函数创建
	closed  bool                      //资源池是否关闭
}

//创建一个资源池
// fn是创建新资源的函数；还有一个size是指定资源池的大小
func NewPool(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size太小了")
	}

	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil

}

//从资源池中取出一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:获取一个共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:生成一个新的资源")
		return p.factory()
	}
}

//关闭资源池，释放资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock() //保证资源关闭正常进行
	if p.closed {
		return
	}
	p.closed = true

	//关闭了通道，就不能再写入
	close(p.res)

	//依次关闭通道里所有资源
	for r := range p.res {
		r.Close()
	}
}

// 释放资源
//释放资源本质上就会把资源再发送到缓冲通道中
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock() //保证该操作和close方法是安全执行的

	//如果资源都关闭了，就不需要释放
	if p.closed {
		r.Close()
		return
	}

	//判断资源是否可以放回资源池中
	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		//如果不能放回，就直接关闭
		log.Println("资源池满了，释放该资源")
		r.Close()
	}
}
