package middleaware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)
var Logger *zap.SugaredLogger

func init(){
	//使用new需要参数encoder、writesyncer、level
	encoder:=zapcore.NewConsoleEncoder(generateEncoderConfig()) //使用默认encoder配置

	//fileHandler,err:=os.OpenFile("..\\zap.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	//if err!=nil{
	//	fmt.Println("Fail to open log file!")
	//}
	//writeSyncer:=zapcore.AddSync(fileHandler)

	//使用lumberjack进行日志切割
	writeSyncer:=getLogWriter()

	logt:=zap.New(zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel),zap.AddCaller())

	Logger =logt.Sugar()
}
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
func generateEncoderConfig() zapcore.EncoderConfig {
	config:=zap.NewProductionEncoderConfig()
	//config.EncodeTime=zapcore.ISO8601TimeEncoder
	config.EncodeTime=func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.EncodeLevel=zapcore.CapitalLevelEncoder
	return config
}

//日志中间件
func GenerateLog()gin.HandlerFunc  {
	return func(context *gin.Context) {
		start:=time.Now()
		context.Next()
		//时间记录
		runTime:=time.Since(start).Seconds()

		urlPath:=context.FullPath()
		useIP:=context.ClientIP()
		method:=context.Request.Method
		statusCode:=context.Writer.Status()
		Logger.Infof("%s\t%d\t%s\t%s\t%.3f",method,statusCode,urlPath,useIP,runTime)
	}
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}