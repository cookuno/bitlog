package bitlog

import (
	"log"
	"io"
	"sync"
)

type BitLog struct  {
	chain []chan DataPkg
	lk *sync.Mutex
}

func New() *BitLog {
	bitLog := &BitLog{chain:make([]chan DataPkg, 8), lk:new(sync.Mutex)}
	return bitLog
}

func (self *BitLog) AddOutput(level int, f Formatter, w Writer)  {
	self.lk.Lock()
	defer self.lk.Unlock()
	p := output{level:level, formatter:f, writer:w, logger: self}
	c := make(chan DataPkg)
	self.chain = append(self.chain, c)
	go func(p *output, c chan DataPkg) {
		go p.on(c)
	}(&p, c)
}

func (self *BitLog) Close() {
	for _, c := range self.chain {
		close(c)
	}
}

func (self *BitLog) Print() (*Printer)  {
	return &Printer{logger: self, dp: DataPkg{}}
}

func (self *BitLog) send(dp DataPkg) {
	for _, c := range self.chain {
		go func(dp DataPkg, c chan DataPkg) {
			c <- dp
		}(dp, c)
	}
}










