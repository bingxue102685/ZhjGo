package conf

var (
	LenStackBuf     = 4096
	LogLevel        string
	LogPath         string
	LogFlag         int
	ConsolePort     int
	ConsolePrompt   string = "ZhjGo# "
	ProfilePath     string
	ListenAddr      string
	ConnAddrs       []string
	PendingWriteNum int
)
