package middleware

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)
func init(){
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)
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

