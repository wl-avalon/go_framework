package log

type JdbLogger struct {
	logPath logPathList
}

type logPathList struct {
	traceLogPath   string
	debugLogPath   string
	noticeLogPath  string
	warningLogPath string
	fatalLogPath   string
	requestLogPath string
}
