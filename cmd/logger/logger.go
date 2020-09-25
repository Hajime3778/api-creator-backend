package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingSetting ログの設定をします。
func LoggingSetting() {

	// 日付単位でログファイルを作成する。
	day := time.Now()
	const layout = "2006-01-02"
	filePath := "./log/" + day.Format(layout) + ".log"

	// ファイルに権限を付与する。
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}

	// 標準出力とファイルの両方を出力先に設定する。
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
	log.SetOutput(gin.DefaultWriter)
}
