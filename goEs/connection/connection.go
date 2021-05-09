package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"go_es/constant"

	"github.com/olivere/elastic/v7"
)

var ES_CLIENT *elastic.Client

func GetConnection() *elastic.Client {
	url := fmt.Sprintf("%s:%d", constant.ES_URL, constant.ES_PORT)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(10*time.Second))
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	ES_CLIENT = client
	return ES_CLIENT
}
