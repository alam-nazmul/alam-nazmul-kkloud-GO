package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	localPath := "/home/ubuntu/daily_photo_backup/data/"
	logFile := fmt.Sprint("/home/ubuntu/daily_photo_backup/logs/metadata_remove_%s_d.log", time.Now().Format("1992-12-25"), os.Getgid()

}
