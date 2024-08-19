package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetOutput(os.Stdout)
}
