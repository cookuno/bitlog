package bitlog

type Formatter interface  {
	Format(*DataPkg) ([]byte, error)
}

const (
	COLOR_DEFAULT 	= "\033[0m"
	COLOR_RED 	= "\033[31m"
	COLOR_GREEN 	= "\033[32m"
	COLOR_YELLOW	= "\033[33m"
	COLOR_BLUE 	= "\033[34m"
	COLOR_ATTR	= "\033[36m"


)
