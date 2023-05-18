# 目标配置

## 测试数据
```json
{
  "name": "ops",
  "url": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=",
  "secret": "",
  "rule": [
    {
      "enabled": true,
      "label_type": "GROUP",
      "label": "customer_city",
      "value": [
        "济南",
        "济宁",
        "德州",
        "青岛"
      ],
      "active": "NEWS",
      "mention": {
        "all": false,
        "mobiles": [
          ""
        ]
      }
    },
    {
      "enabled": true,
      "label_type": "COMMON",
      "label": "status",
      "value": [
        "resolved"
      ],
      "active": "STATUS_TRANS",
      "mention": {
        "all": false,
        "mobiles": [
          ""
        ]
      }
    }
  ]
}
```