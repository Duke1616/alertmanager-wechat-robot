# 报警规则模块
## 示例1
```json
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
}
```

## 示例2
```json
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
```