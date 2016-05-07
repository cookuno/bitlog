package bitlog

import (
	"errors"
	"fmt"
)

type TextLineFormatter struct  {

}

func (self TextLineFormatter) Format(dataPkg *DataPkg) ([]byte, error)  {
	if dataPkg == nil {
		return nil, errors.New("dataPkg is nil")
	}
	textLine := ""
	if dataPkg.Level > 0 {
		level := ""
		if dataPkg.Level == LEVEL_DEBUG {
			level = COLOR_BLUE + dataPkg.levelName() + COLOR_DEFAULT
		} else if dataPkg.Level == LEVEL_INFO {
			level = COLOR_GREEN + dataPkg.levelName() + COLOR_DEFAULT
		} else if dataPkg.Level == LEVEL_WARN {
			level = COLOR_YELLOW + dataPkg.levelName() + COLOR_DEFAULT
		} else if dataPkg.Level == LEVEL_ERROR {
			level = COLOR_RED + dataPkg.levelName() + COLOR_DEFAULT
		} else if dataPkg.Level == LEVEL_PANIC {
			level = COLOR_RED + dataPkg.levelName() + COLOR_DEFAULT
		} else if dataPkg.Level == LEVEL_FATAL {
			level = COLOR_RED + dataPkg.levelName() + COLOR_DEFAULT
		}
		level = "LEVEL=" + level
		textLine = textLine + level + " | "
	}
	if dataPkg.Time != nil || dataPkg.Time.IsZero() {
		textLine = textLine + "TIME=" + dataPkg.Time.String() + " | "
	}
	textLine = textLine + " " + dataPkg.Message + " | "
	if dataPkg.Attr != nil {
		attrLine := ""
		for _, attr := range dataPkg.Attr {
			key := COLOR_ATTR + fmt.Sprint(attr.Key) + COLOR_DEFAULT
			val := fmt.Sprint(attr.Val)
			attrLine = attrLine + key + "=" + val + " ; "
		}
		if attrLine != "" {
			textLine = textLine + " " + attrLine + " | "
		}
	}
	if dataPkg.Position == "" {
		textLine = textLine + "LINE=" + dataPkg.Position + " "
	}
	return []byte(textLine + "\n"), nil
}
