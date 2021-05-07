### 功能描述
推广员新增

### 请求说明
> 请求方式：POST<br>
请求URL : api/proxy/act/mapi/promoter/add <br>
请求格式：application/json <br>
开发者： 陆伟杰

### 请求参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
| group_id | integer | Yes | group_id的注释 |
| list | array | Yes | 列表 |
| ↳task_id | integer | Yes | task_id的注释 |
| ↳task_kind | integer | Yes | task_kind的注释 |
| ↳task_title | string | Yes | task_title的注释 |
| ↳list | array | Yes | 列表 |
| ↳↳reward_value | string | Yes | reward_value的注释 |
| ↳↳id | integer | Yes | ID号 |
| ↳↳reward_rule | integer | Yes | reward_rule的注释 |
| ↳↳reward_type | integer | Yes | reward_type的注释 |

### 返回参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
| code | integer | Yes | 编码 |
| data | object | Yes | 数据 |
| ↳list | array | Yes | 列表 |
| ↳↳group_id | integer | Yes | group_id的注释 |
| ↳↳group_title | string | Yes | group_title的注释 |
| ↳↳group_status | integer | Yes | group_status的注释 |
| ↳↳update_time | integer | Yes | update_time的注释 |
| msg | string | Yes | 信息 |
| trace_id | string | Yes | trace_id的注释 |

### 错误状态码
| 状态码 | 说明 |
| ----- | ---- |
| 0     | 正常 |
| 非0   | 发生错误 |

### 请求示例
```shell
curl -X POST 'http://dev-api-chengjinsheng.mingqijia.com/api/proxy/act/mapi/promoter/add' \
-H 'Content-Type: application/json' \
-H 'branchname: dev' \
-d '
{
	"group_id": 1,
	"list": [
		{
			"list": [
				{
					"id": 1,
					"reward_rule": 11,
					"reward_type": 2,
					"reward_value": "5"
				},
				{
					"id": 2,
					"reward_rule": 21,
					"reward_type": 2,
					"reward_value": "10"
				},
				{
					"id": 3,
					"reward_rule": 22,
					"reward_type": 2,
					"reward_value": "15"
				},
				{
					"id": 4,
					"reward_rule": 23,
					"reward_type": 3,
					"reward_value": "ts001,ts002,ts003"
				}
			],
			"task_id": 1,
			"task_kind": 11,
			"task_title": "饮食记录奖励"
		},
		{
			"list": [
				{
					"id": 5,
					"reward_rule": 11,
					"reward_type": 2,
					"reward_value": "5"
				},
				{
					"id": 6,
					"reward_rule": 21,
					"reward_type": 2,
					"reward_value": "10"
				},
				{
					"id": 7,
					"reward_rule": 22,
					"reward_type": 2,
					"reward_value": "15"
				},
				{
					"id": 8,
					"reward_rule": 23,
					"reward_type": 3,
					"reward_value": "ts001,ts002,ts003"
				}
			],
			"task_id": 2,
			"task_kind": 12,
			"task_title": "体重记录奖励"
		}
	]
}
'
```

### 返回结果示例
```json
{
	"code": 0,
	"data": {
		"list": [
			{
				"group_id": 1,
				"group_status": 0,
				"group_title": "每日健康打卡",
				"update_time": 1.614675531e+09
			},
			{
				"group_id": 2,
				"group_status": 1,
				"group_title": "每日问答",
				"update_time": 1.614675531e+09
			},
			{
				"group_id": 3,
				"group_status": 0,
				"group_title": "填写个人健康数据",
				"update_time": 1.614675531e+09
			}
		]
	},
	"msg": "success",
	"trace_id": "83a1ec395a1775cb62a1d8b89366727e"
}
```
