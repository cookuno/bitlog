package bitlog

import (
	"time"
	"runtime"
	"fmt"
)

var (
	reset 	= "\033[0m"
	red 	= "\033[31m"
	yellow	= "\033[33m"
	blue 	= "\033[34m"
	green 	= "\033[32m"
)

type Formatter interface  {
	Format(level string, v interface{}) string
}

type DefaultFormatter struct  {
	DateTimeFormat string
	Position bool
}

func (self DefaultFormatter) Format(level string, v interface{}) string {
	output :=  level + "|"
	if self.DateTimeFormat != "" {
		dateTime := time.Now().Format(self.DateTimeFormat)
		output = output + "|" + dateTime
	}
	if self.Position {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		output = output + "|" + file + ":" + fmt.Sprint(line)
	}
	output = output + "|" + fmt.Sprint(v)
	if level == LEVEL_ERROR {
		output = red + output + reset
	} else if level == LEVEL_WARN {
		output = yellow + output + reset
	} else if level == LEVEL_DEBUG {
		output = blue + output + reset
	} else if level == LEVEL_TRACE {
		output = green + output + reset
	}
	return output
}
