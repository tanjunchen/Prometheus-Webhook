package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"

	"webhook.prometheus/model"
	"webhook.prometheus/notifier"
)
var (
	h           bool
	defaultRobot string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "default", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
}

func main()  {
	flag.Parse()

	if h{
		flag.Parse()
		return
	}

	router := gin.Default()
	router.POST("/webhook", func(context *gin.Context) {
		var notification model.Notification

		err := context.BindJSON(&notification)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		context.JSON(http.StatusOK, gin.H{"message": "send to dingTalk successful!"})
	})

	router.Run()
}