### 功能描述
%s

### 请求说明
> 请求方式：POST<br>
请求URL ：%s <br>
请求格式：application/json <br>
开发者： 陆伟杰

### 请求参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
%s

### 返回参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
%s

### 错误状态码
| 状态码 | 说明 |
| ----- | ---- |
| 0     | 正常 |
| 非0   | 发生错误 |

### 请求示例
```shell
curl -X POST 'http://dev-manage-chenyinfei.mingqijia.com/%s' \
-H 'Content-Type: application/json' \
-H 'branchname: dev' \
-d '
%s
'

```

### 返回结果示例
```json
%s
```