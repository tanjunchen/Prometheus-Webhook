# Note

The Key Words of dingtalkRobot must be consistent with msg from prometheus.

The Debug manner for webhook.

curl -l -H "Content-type: application/json" -X POST -d 
'{"msgtype": "markdown","markdown": {"title":"Prometheus告警信息","text": "#### 监控指标\n> 监控描述信息\n\n> ###### 告警时间 \n"},"at": {"isAtAll": false}}' 
https://oapi.dingtalk.com/robot/send?access_token=xxxxxx


thanks https://github.com/yunlzheng/prometheus-book
