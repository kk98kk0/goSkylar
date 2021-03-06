package logUtil

import (
	"log"
	"os"
	"time"
)

func LogSet() { //解析参数付给logF
	outfile, err := os.OpenFile("log/"+CurrentDate()+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名
}

// 获取当天的日期
func CurrentDate() string {
	return DateToStr(time.Now().Unix())
}

func DateToStr(intTime int64) string {
	timeLayout := "2006-01-02"
	dataTimeStr := time.Unix(intTime, 0).Format(timeLayout)
	return dataTimeStr
}
