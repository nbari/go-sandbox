package main

import (
	"fmt"
	"time"
)

type Producer struct {
	chs []chan<- int
}

func (p *Producer) Register(ls ...*Listener) {
	for _, l := range ls {
		p.chs = append(p.chs, l.ch)
	}
}

func (p *Producer) Fanout(i int) {
	for _, ch := range p.chs {
		ch := ch
		go func(i int) {
			ch <- i
		}(i)
	}
}

type Listener struct {
	ch chan int
}

func (l *Listener) Listen(prefix string) {
	go func(p string) {
		for i := range l.ch {
			fmt.Println(p, i)
		}
	}(prefix)
}

func main() {
	l1 := &Listener{make(chan int, 0)}
	l1.Listen("listener 1:")

	l2 := &Listener{make(chan int, 0)}
	l2.Listen("listener 2:")

	p := &Producer{make([]chan<- int, 0)}
	p.Register(l1, l2)
	p.Fanout(5)

	time.Sleep(2 * time.Second)
	fmt.Println("exit")

}
