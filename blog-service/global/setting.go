package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)