---
title: Paper v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.11"

---

# Paper

> v1.0.0

# 鉴权

## GET 获取用户信息

GET /api/v1/auth/userinfo

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 普通用户注册

POST /api/v1/auth/register

普通用户注册账号，用户信息写入mysql数据库

> Body 请求参数

```yaml
username: string
email: string
telephone: string
organization: string
address: string
postcode: string
password: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |用户注册的用户名，需要唯一|
|» email|body|string| 是 |用户注册的邮箱，需要唯一|
|» telephone|body|string| 是 |用户注册的手机号，需要唯一，是唯一登录标识|
|» organization|body|string| 否 |用户注册的组织|
|» address|body|string| 否 |用户注册的地址|
|» postcode|body|string| 否 |用户注册的邮编|
|» password|body|string| 是 |用户注册设置密码|

> 返回示例

> 用户注册成功

```json
{
  "data": {},
  "status": 200,
  "error": {
    "msg": "",
    "code": ""
  }
}
```

> 用户注册失败（注册请求格式有误）

```json
{
  "data": {},
  "status": 400,
  "error": {
    "msg": "user register message format error",
    "code": 6
  }
}
```

> 用户注册失败（信息重复）

```json
{
  "data": {},
  "status": 401,
  "error": {
    "msg": "user eamil is duplicated",
    "code": 7
  }
}
```

> 用户注册服务器内部错误

```json
{
  "data": {},
  "status": 500,
  "error": {
    "msg": "",
    "code": 5
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|用户注册成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|用户注册失败（注册请求格式有误）|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|用户注册失败（信息重复）|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|用户注册服务器内部错误|Inline|

### 返回数据结构

状态码 **201**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

## GET 获取所有用户

GET /api/v1/users

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|idx|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 普通用户登录

POST /api/v1/auth/login

普通用户通过该接口登录，获取token，之后请求附带token

> Body 请求参数

```json
{
  "telephone": "18019134181",
  "password": "hxc123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Origin|header|string| 否 |none|
|body|body|object| 否 |none|
|» telephone|body|string| 是 |普通用户登录手机号|
|» password|body|string| 是 |普通用户登录密码|

> 返回示例

> 用户成功

```json
{
  "data": {
    "token": "qulaborusedamet"
  },
  "status": 200,
  "error": {
    "msg": "",
    "code": ""
  }
}
```

> 用户登录请求格式有误

```json
{
  "data": {},
  "status": 404,
  "error": {
    "msg": "user is not exist",
    "code": 10
  }
}
```

> 用户登录密码错误

```json
{
  "data": {},
  "status": 401,
  "error": {
    "msg": "user password error",
    "code": 11
  }
}
```

> 用户登录用户不存在

```json
{
  "data": {},
  "status": 404,
  "error": {
    "msg": "user is not exist",
    "code": 10
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|用户成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|用户登录请求格式有误|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|用户登录密码错误|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|用户登录用户不存在|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|用户服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» Token|string|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **404**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|» status|integer|true|none||none|
|» error|object|true|none||none|
|»» msg|string|true|none||none|
|»» code|integer|true|none||none|

# 鉴权/管理员操作权限

## PUT 修改用户权限

PUT /api/v1/mode

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|
|mode|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 修改用户信息

PUT /api/v1/users

> Body 请求参数

```yaml
id: string
username: string
organization: string
address: string
postcode: string
password: string
email: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|
|» username|body|string| 是 |none|
|» organization|body|string| 是 |none|
|» address|body|string| 是 |none|
|» postcode|body|string| 是 |none|
|» password|body|string| 是 |none|
|» email|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 投稿

## DELETE 删除投稿

DELETE /api/v1/submission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 修改投稿

PUT /api/v1/submission

> Body 请求参数

```yaml
id: string
author: string
cauthor: string
sauthor: string
tauthor: string
fauthor: string
topic: 0
keyword: string
file: string
name: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|
|» author|body|string| 否 |none|
|» cauthor|body|string| 否 |none|
|» sauthor|body|string| 否 |none|
|» tauthor|body|string| 否 |none|
|» fauthor|body|string| 否 |none|
|» topic|body|integer| 否 |none|
|» keyword|body|string| 否 |none|
|» file|body|string(binary)| 否 |none|
|» name|body|string| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建投稿

POST /api/v1/submission

> Body 请求参数

```yaml
name: string
author: string
cauthor: string
sauthor: string
tauthor: string
fauthor: string
topic: 0
keyword: string
file: string
abstract: string
organization: string
address: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» name|body|string| 是 |none|
|» author|body|string| 是 |none|
|» cauthor|body|string| 是 |none|
|» sauthor|body|string| 否 |none|
|» tauthor|body|string| 否 |none|
|» fauthor|body|string| 否 |none|
|» topic|body|integer| 是 |none|
|» keyword|body|string| 是 |none|
|» file|body|string(binary)| 是 |none|
|» abstract|body|string| 是 |none|
|» organization|body|string| 是 |none|
|» address|body|string| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取投稿

GET /api/v1/submission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取投稿对应所有审稿

GET /api/v1/contributor-papers-review

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 根据paper名称获取投稿

GET /api/v1/searchsubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|name|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取全部投稿分页

GET /api/v1/submission/admin/paging

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|idx|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取用户投稿

GET /api/v1/submission/user

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 下载投稿对应论文

GET /api/v1/submission/paper

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取用户投稿分页

GET /api/v1/submission/user/paging

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|idx|query|string| 是 |none|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 投稿/管理员操作投稿

## GET 接受论文投稿

GET /api/v1/approvesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 拒绝论文投稿

GET /api/v1/refusesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 评审/评审人

## GET 获取审稿人

GET /api/v1/reviewer

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建审稿人

POST /api/v1/reviewer

> Body 请求参数

```json
{
  "field1": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|
|body|body|object| 否 |none|
|» field1|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除审稿人

DELETE /api/v1/reviewer

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取审稿人对应的全部论文

GET /api/v1/reviewer-papers-review

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 评审/分配评审文章

## DELETE 删除论文评审分配

DELETE /api/v1/reviewersbm

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取某条reviewersbm

GET /api/v1/reviewersbm

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 新建论文评审分配

POST /api/v1/reviewersbm

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|rid|query|string| 是 |none|
|sid|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取全部论文评审分配

GET /api/v1/reviewersbms

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 评审/评审

## POST 创建审稿评论

POST /api/v1/reviewcmt

> Body 请求参数

```yaml
id: string
comment: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|
|» comment|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Issue

## GET 获取最新出版issue

GET /api/v1/latestissue

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Issue/issue操作

## DELETE 删除issue

DELETE /api/v1/issue

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 发表issue

PUT /api/v1/issue

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取issue

GET /api/v1/issue

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建新期

POST /api/v1/issue

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取全部issues

GET /api/v1/allissue

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Issue/投稿-issue分配

## DELETE 从issue中删除投稿

DELETE /api/v1/issuesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|sid|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 修改投稿对应issue

PUT /api/v1/issuesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|sid|query|string| 是 |none|
|iid|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 分配投稿至issue

POST /api/v1/issuesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|iid|query|string| 是 |none|
|sid|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取issue全部投稿

GET /api/v1/issuesubmission

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|iid|query|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

