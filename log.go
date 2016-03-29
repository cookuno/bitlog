package bitlog

import (
	"log"
	"io"
)

const (
	LEVEL_ERROR 		= 0
	LEVEL_WARN 		= 1
	LEVEL_INFO 		= 2
	LEVEL_DEBUG 		= 3
	LEVEL_TRACE 		= 4
	LEVEL_ERROR_NAME 	= "ERROR"
	LEVEL_WARN_NAME 	= "WARN"
	LEVEL_INFO_NAME 	= "INFO"
	LEVEL_DEBUG_NAME 	= "DEBUG"
	LEVEL_TRACE_NAME 	= "TRACE"
)

type BitLog struct  {
	level 		int
	logger 		*log.Logger
	pipe 		chan string
	formatter 	Formatter
}

func New(level int, out io.Writer) *BitLog {
	bitLog := &BitLog{}

	return bitLog
}

/**
 * format v to string, send it into chan pipe
 */
func (self *BitLog) doLog(level string, v interface{}) {
	s := self.formatter.Format(level, v)
	go func(v string, bitLog *BitLog) {
		bitLog.pipe <- v
	}(s, self)
}

func (self *BitLog) checkLevel(targetLevel int) bool {
	if self.level <= targetLevel {
		return true
	}
	return false
}

func (self *BitLog) Error(v interface{}) {
	if !self.checkLevel(LEVEL_ERROR) {
		return
	}
	self.doLog(LEVEL_ERROR_NAME, v)
}


func (self *BitLog) Warn(v interface{}) {
	if !self.checkLevel(LEVEL_WARN) {
		return
	}
	self.doLog(LEVEL_WARN_NAME, v)
}

func (self *BitLog) Info(v interface{}) {
	if !self.checkLevel(LEVEL_INFO) {
		return
	}
	self.doLog(LEVEL_INFO_NAME, v)
}

func (self *BitLog) Debug(v interface{}) {
	if !self.checkLevel(LEVEL_DEBUG) {
		return
	}
	self.doLog(LEVEL_DEBUG_NAME, v)
}

func (self *BitLog) Trace(v interface{}) {
	if !self.checkLevel(LEVEL_TRACE) {
		return
	}
	self.doLog(LEVEL_TRACE_NAME, v)
}


