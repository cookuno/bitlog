package bitlog

import (
	"time"
)

type DataPkg struct {
	Time 		*time.Time
	Position 	string
	Level 		int
	Attr		*Attribute
	Message		string
}

func (self *DataPkg) levelName() string {
	if self.Level == LEVEL_DEBUG {
		return LEVEL_DEBUG_NAME
	} else if self.Level == LEVEL_INFO {
		return LEVEL_INFO_NAME
	} else if self.Level == LEVEL_WARN {
		return LEVEL_WARN_NAME
	} else if self.Level == LEVEL_ERROR {
		return LEVEL_ERROR_NAME
	} else if self.Level == LEVEL_PANIC {
		return LEVEL_PANIC_NAME
	} else if self.Level == LEVEL_FATAL {
		return LEVEL_FATAL_NAME
	} else {
		return "N/A"
	}
}

type Attribute []AttributeEntry

type AttributeEntry struct  {
	Key interface{}
	Val interface{}
}
