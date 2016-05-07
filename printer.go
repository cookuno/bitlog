package bitlog

import (
	"fmt"
	"runtime"
	"os"
)

type Printer struct  {
	logger *BitLog
	dp DataPkg
}

func (self *Printer) Attr(attr Attribute) *Printer {
	self.dp.Attr = attr
	return self
}

func (self *Printer) LineNo() *Printer {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	self.dp.Position = file + ":" + fmt.Sprint(line)
	return self
}

func (self *Printer) Debug(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_DEBUG
	self.logger.send(self.dp)
}

func (self *Printer) DebugF(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_DEBUG
	self.logger.send(self.dp)
}

func (self *Printer) Info(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_INFO
	self.logger.send(self.dp)
}

func (self *Printer) InfoF(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_INFO
	self.logger.send(self.dp)
}

func (self *Printer) Warn(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_WARN
	self.logger.send(self.dp)
}

func (self *Printer) WarnF(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_WARN
	self.logger.send(self.dp)
}

func (self *Printer) Error(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_ERROR
	self.logger.send(self.dp)
}

func (self *Printer) ErrorF(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_ERROR
	self.logger.send(self.dp)
}

func (self *Printer) Panic(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_PANIC
	self.logger.send(self.dp)
	panic(msg)
}

func (self *Printer) PanicF(format string, v ...interface{})  {
	msg := fmt.Sprint(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_PANIC
	self.logger.send(self.dp)
	panic(msg)
}

func (self *Printer) Fatal(v ...interface{})  {
	msg := fmt.Sprint(v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_FATAL
	self.logger.send(self.dp)
	os.Exit(1)
}

func (self *Printer) FatalF(format string, v ...interface{})  {
	msg := fmt.Sprintf(format, v)
	self.dp.Message = msg
	self.dp.Level = LEVEL_FATAL
	self.logger.send(self.dp)
	os.Exit(1)
}