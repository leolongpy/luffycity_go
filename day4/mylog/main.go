package mylog

const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	CRATAL
)

func getLevelStr(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "CRATAL"
	default:
		return "DEBUG"
	}
}
