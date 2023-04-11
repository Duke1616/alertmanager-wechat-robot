# 企业微信群组robot配置

## 测试数据
```json
{
    "name": "ops",
    "url": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=",
    "secret": "",
    "rule": [
        {
            "enabled": true,
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