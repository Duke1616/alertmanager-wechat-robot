# 接收报警信息

## 测试数据
```json
{
	"receiver": "wechatbot.hook",
	"status": "firing",
	"alerts": [{
		"status": "firing",
		"labels": {
			"alertgroup": "vmagent",
			"alertname": "模拟测试",
			"customer_city": "济南",
			"customer_name": "XXX",
            "customer_province": "山东",
			"instance": "XXX:8429",
			"job": "vmagent",
			"severity": "warning",
			"source": "vmagent",
			"status": "resolved"
		},
		"annotations": {
			"description": "VPN恢复连接",
			"summary": "市：济南 节点: XXX"
		},
		"startsAt": "2023-04-11T07:37:10.853043175Z",
		"endsAt": "0001-01-01T00:00:00Z"
	}, {
		"status": "firing",
		"labels": {
			"alertgroup": "vmagent",
			"alertname": "模拟测试",
			"customer_city": "济南",
			"customer_name": "XXX",
            "customer_province": "山东",
			"instance": "10.31.0.18:8429",
			"job": "vmagent",
			"severity": "warning",
			"source": "vmagent",
			"status": "resolved"
		},
		"annotations": {
			"description": "VPN恢复连接",
			"summary": "市：济南 稍点单位: XXX"
		},
		"startsAt": "2023-04-11T07:37:10.853043175Z",
		"endsAt": "0001-01-01T00:00:00Z"
	}],
	"groupLabels": {
		"severity": "warning",
		"source": "vmagent"
	},
	"commonLabels": {
		"alertgroup": "vmagent",
		"alertname": "模拟测试",
		"job": "vmagent",
		"severity": "warning",
		"source": "vmagent",
		"status": "resolved"
	},
	"commonAnnotations": {
		"description": "VPN恢复连接"
	},
	"externalURL": "http://alertmanager-6b9v9hsc:9093",
	"version": "4",
	"truncatedAlerts": 0
}
```