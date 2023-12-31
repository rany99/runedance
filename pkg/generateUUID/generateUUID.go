package generateUUID

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
)

func MyRequestid() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		LoggerInit()
		requestId := uuid.New()

		c.Set("request_id", requestId.String())

	}
}

var logger = &lumberjack.Logger{
	Filename:   "../../log/log.txt",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     28, //days
}

func LoggerInit() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:03:04",

		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名
			base := path.Base(frame.File)
			pre := path.Dir(frame.File)
			pre = path.Base(pre)
			realfilename := pre + "/" + base
			return frame.Function, realfilename
		},
	})
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logger)
	logrus.SetOutput(fileAndStdoutWriter)
	//设置最低loglevel

	logrus.SetLevel(logrus.InfoLevel)

}
