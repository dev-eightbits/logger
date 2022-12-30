package logger

type LoggerConfig struct {
	Level           string `default:"INFO"`
	Format          string `default:"TEXT"`
	Output          string `default:"Stdout"`
	TimestampFormat string `default:"2006-01-02 15:04:05"`
	FullTimestamp   bool   `default:"true"`
	DisableColor    bool   `default:"true"`
	ReportCaller    bool   `default:"true"`
	PrettyPrint     bool   `default:"true"`
}
