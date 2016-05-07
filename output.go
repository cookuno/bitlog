package bitlog

const (
	LEVEL_FATAL		= 1
	LEVEL_PANIC		= 2
	LEVEL_ERROR 		= 3
	LEVEL_WARN 		= 4
	LEVEL_INFO 		= 5
	LEVEL_DEBUG 		= 6
	LEVEL_FATAL_NAME 	= "FATAL"
	LEVEL_PANIC_NAME 	= "PANIC"
	LEVEL_ERROR_NAME 	= "ERROR"
	LEVEL_WARN_NAME 	= "WARN"
	LEVEL_INFO_NAME 	= "INFO"
	LEVEL_DEBUG_NAME 	= "DEBUG"
)

type output struct  {
	level 		int
	writer 		Writer
	formatter 	Formatter
	logger 		*BitLog
}

func (self *output) on(c chan DataPkg) {
	var dp DataPkg = nil
	dp <- c
	if dp == nil {
		return
	}
	if self.level == 0 || self.level < dp.Level {
		return
	}
	outputData, fmtErr := self.formatter.Format(&dp)
	if fmtErr != nil {
		self.logger.Print().Panic(fmtErr)
	}
	_, outputErr := self.writer.Write(outputData)
	if outputErr != nil {
		self.logger.Print().Panic(outputErr)
	}
}
