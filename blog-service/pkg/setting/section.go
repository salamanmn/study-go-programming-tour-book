package setting

import "time"

type ServerSettingS struct {
	RunMode     string
	HttpPort    string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
}

type DatabaseSettingS struct {
  DBType       string 
  UserName string
  Password string
  Host string
  DBName string
  TablePrefix  string
  Charset string
  ParseTime bool
  MaxIdleConns int
  MaxOpenConns int
}
