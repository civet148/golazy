package gen

const (
	internal         = "internal/"
	typesPacket      = "types"
	configPacket     = "config"
	svcPacket        = "svc"
	handlerPacket    = "handler"
	logicPacket      = "logic"
	middlewarePacket = "middleware"
	configDir        = internal + configPacket
	contextDir       = internal + svcPacket
	handlerDir       = internal + handlerPacket
	logicDir         = internal + logicPacket
	middlewareDir    = internal + middlewarePacket
	typesDir         = internal + typesPacket
	groupProperty    = "group"
)

const (
	// ProjectName the const value of zero
	ProjectName = "zero"
	// ProjectOpenSourceURL the github url of go-zero
	ProjectOpenSourceURL = "github.com/zeromicro/go-zero"
	ProjectLogURL        = "github.com/civet148/log"
	ProjectDatabaseURL   = "github.com/civet148/sqlca"
	ProjectGinURL        = "github.com/gin-gonic/gin"
	// OsWindows represents os windows
	OsWindows = "windows"
	// OsMac represents os mac
	OsMac = "darwin"
	// OsLinux represents os linux
	OsLinux = "linux"
	// OsJs represents os js
	OsJs = "js"
	// OsIOS represents os ios
	OsIOS = "ios"
)

