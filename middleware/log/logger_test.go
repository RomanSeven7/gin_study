package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)
func init(){
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
	rotateFileHook, _ := NewRotateFileHook(*Generate(logrus.DebugLevel))
	rotateFileHook2, _ := NewRotateFileHook(*Generate(logrus.InfoLevel))
	rotateFileHook3, _ := NewRotateFileHook(*Generate(logrus.WarnLevel))
	rotateFileHook4, _ := NewRotateFileHook(*Generate(logrus.ErrorLevel))
	logrus.AddHook(rotateFileHook)
	logrus.AddHook(rotateFileHook2)
	logrus.AddHook(rotateFileHook3)
	logrus.AddHook(rotateFileHook4)

}

func TestPrint(t *testing.T) {
	for i := 0; i <10000; i++ {
		//time.Sleep(time.Second)
		logrus.Debug("debug 111")
		logrus.Info("info 111")
		logrus.Warn("warn 111")
		logrus.Error("Error 111")
	}

}

func Generate(level logrus.Level)*RotateFileConfig{
	 rotateFileConfig:=  RotateFileConfig{
		Filename: fmt.Sprintf("./log/%s-%s.log",level.String(),time.Now().Format(time.RFC3339)),
		MaxSize: 10,
		MaxBackups: 10,
		MaxAge: 7,
		Level: level,
		Formatter: &logrus.JSONFormatter{},
	}
	writer, _ := rotatelogs.New(
		rotateFileConfig.Filename+".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(rotateFileConfig.Filename),
		rotatelogs.WithMaxAge(time.Duration(2)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(5)*time.Second),
	)
	logrus.SetOutput(writer)
	return &rotateFileConfig
}