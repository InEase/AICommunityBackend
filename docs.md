---
title: AI Community v1.0.0
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
generator: "@tarslib/widdershins v4.0.4"

---

# AI Community

> v1.0.0

# 用户管理

## DELETE 删除用户

DELETE /user/info

后台验证，要删除的用户是否是本人

> 返回示例

> 权限验证失败

```json
{
  "code": "1005",
  "data": null,
  "msg": "无权限"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|权限验证失败|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|null|true|none|none|
|» msg|string|true|none|none|

## GET 获取用户信息

GET /user/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|cookie|string|false|none|
|userid|query|string|false|如果不给出，则查询当前登录的用户信息|

> 返回示例

> 未鉴权

```json
{
  "code": 1000,
  "data": null,
  "msg": "无权限"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|未鉴权|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|number|true|none|none|
|» data|null|true|none|none|
|» msg|string|true|none|none|

## POST 更改用户信息

POST /user/info

默认修改本人信息

> Body 请求参数

```yaml
name: 小明
gender: 男
school: **大学
avatar: string
desc: 简介

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» name|body|string|false|用户名称|
|» gender|body|string|false|性别|
|» school|body|string|false|学校|
|» avatar|body|string|false|头像url，必须满足正则，如果不传后端会自动生成一个头像作为用户头像|
|» desc|body|string|false|简介|

> 返回示例

> 成功

```json
{
  "code": 0,
  "field2": {
    "name": "黎娟",
    "userid": "7420250294029280",
    "gender": "男",
    "avatar": "http://lotxcbb.pw/yjkmdepek",
    "desc": "因整话派战何无度矿型却压白。生干然合酸所价切值记元车。社实程来程型切风温解处没就相六。织多事强际了年照家采又便。专酸空声争节受什热效于百。切前在二在位见属同开风议际发构消。",
    "alt": "http://gnmdxjzm.sd/rmrj",
    "school": "于静",
    "status": "ad cupidatat nostrud deserunt"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|0|
|» data|[%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF](#schema%e7%94%a8%e6%88%b7%e4%bf%a1%e6%81%af)|true|none|none|
|»» name|string|true|none|none|
|»» status|string|true|none|none|
|»» userid|integer|true|none|保证非负|
|»» gender|string|true|none|none|
|»» avatar|string|true|none|none|
|»» desc|string|true|none|none|
|»» school|string|true|none|none|
|» msg|string|true|none|none|

## POST 用户注册

POST /user/register

> Body 请求参数

```yaml
name: 小明
password: string
gender: 男
school: **大学
avatar: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» name|body|string|true|用户名称|
|» password|body|string|true|none|
|» gender|body|string|true|性别|
|» school|body|string|true|学校|
|» avatar|body|string|false|头像url，必须满足正则，如果不传后端会自动生成一个头像作为用户头像|

> 返回示例

> 参数错误

```json
{
  "code": "1002",
  "data": null,
  "msg": "参数错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|参数错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|错误码可能会有不同，具体看文档|
|» data|null|true|none|none|
|» msg|string|true|none|none|

## POST 用户登录

POST /user/login

> Body 请求参数

```yaml
name: string
password: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» name|body|string|true|none|
|» password|body|string|true|none|

> 返回示例

> 成功

```json
{
  "code": "0",
  "data": {
    "token": "#ZurmEE"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|object|true|none|none|
|»» token|string|true|none|前端需要保存token以鉴权|
|» msg|string|true|none|none|

## GET 搜索用户

GET /user/search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|userid|query|string|false|请求参数任意组合|
|name|query|string|false|none|
|school|query|string|false|none|
|gender|query|string|false|none|
|status|query|string|false|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": [
    {
      "name": "史伟",
      "userid": "8868540737107566",
      "gender": "男",
      "avatar": "http://xviirfcs.as/lswxuwrxh",
      "desc": "农候常算向指叫律比义更气育铁实。团眼做海点此手资己光电府治决容示江。商物性广立根者向众采政众位今门而光称。色保题于育空给即根她属要传图比术他取。",
      "alt": "http://clvrebwkg.cm/wdbsv",
      "school": "八"
    },
    {
      "name": "史伟",
      "userid": "8868540737107566",
      "gender": "男",
      "avatar": "http://xviirfcs.as/lswxuwrxh",
      "desc": "农候常算向指叫律比义更气育铁实。团眼做海点此手资己光电府治决容示江。商物性广立根者向众采政众位今门而光称。色保题于育空给即根她属要传图比术他取。",
      "alt": "http://clvrebwkg.cm/wdbsv",
      "school": "八"
    },
    {
      "name": "史伟",
      "userid": "8868540737107566",
      "gender": "男",
      "avatar": "http://xviirfcs.as/lswxuwrxh",
      "desc": "农候常算向指叫律比义更气育铁实。团眼做海点此手资己光电府治决容示江。商物性广立根者向众采政众位今门而光称。色保题于育空给即根她属要传图比术他取。",
      "alt": "http://clvrebwkg.cm/wdbsv",
      "school": "八"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|string|true|none|none|
|» data|[%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF](#schema%e7%94%a8%e6%88%b7%e4%bf%a1%e6%81%af)|true|none|none|
|»» name|string|true|none|none|
|»» status|string|true|none|none|
|»» userid|integer|true|none|保证非负|
|»» gender|string|true|none|none|
|»» avatar|string|true|none|none|
|»» desc|string|true|none|none|
|»» school|string|true|none|none|
|» msg|string|true|none|none|

# 文章管理

## DELETE 删除文章

DELETE /article/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|articleid|query|string|false|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 更改文章信息

POST /article/info

需要鉴权

> Body 请求参数

```yaml
articleid: string
title: string
body: string
category: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» articleid|body|string|true|none|
|» title|body|string|false|none|
|» body|body|string|false|none|
|» category|body|string|false|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": {
    "articleid": 9,
    "title": "志设干须此花",
    "body": "当直因生林加极许团目接律内维也军。别按亲华则马系两红步研调情厂委起各。得务事形且半实花分五对其做程。",
    "like": 8926646864720792
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|[%E6%96%87%E7%AB%A0%E4%BF%A1%E6%81%AF](#schema%e6%96%87%e7%ab%a0%e4%bf%a1%e6%81%af)|true|none|none|
|»» articleId|number|true|none|保证非负|
|»» title|string|true|none|标题|
|»» body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|»» like|number|true|none|保证非负|
|»» category|number|true|none|none|
|»» category_name|string|true|none|对应分类名|
|» msg|string|true|none|none|

## GET 获取文章信息

GET /article/info

无需鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|articleid|query|string|false|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": {
    "articleid": 3,
    "title": "矿要委入机保元",
    "body": "商则干确表知总万指把作理。四府技动影没八完除物主建务面样天。东强权院么手内志几线消太性人按拉工两。热教引九变己北象交成事改种广建器律。第细圆据容万天必平按社称被将边第器论。",
    "like": -3224000790885832
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|[%E6%96%87%E7%AB%A0%E4%BF%A1%E6%81%AF](#schema%e6%96%87%e7%ab%a0%e4%bf%a1%e6%81%af)|true|none|none|
|»» articleId|number|true|none|保证非负|
|»» title|string|true|none|标题|
|»» body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|»» like|number|true|none|保证非负|
|»» category|number|true|none|none|
|»» category_name|string|true|none|对应分类名|
|» msg|string|true|none|none|

## POST 收藏

POST /article/Favorite

> Body 请求参数

```yaml
articleid: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» articleid|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取文章推荐列表

GET /article/list

返回的是id列表~

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|limit|query|string|false|none|
|offset|query|string|false|none|
|category|query|string|false|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": [
    50,
    57,
    49,
    5,
    30
  ],
  "msg": "consectetur velit dolore"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|[[%E6%96%87%E7%AB%A0%E4%BF%A1%E6%81%AF](#schema%e6%96%87%e7%ab%a0%e4%bf%a1%e6%81%af)]|true|none|none|
|»» articleId|number|true|none|保证非负|
|»» title|string|true|none|标题|
|»» body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|»» like|number|true|none|保证非负|
|»» category|number|true|none|none|
|»» category_name|string|true|none|对应分类名|
|» msg|string|true|none|none|

## POST 发布文章

POST /article/publish

> Body 请求参数

```yaml
title: string
body: string
category: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» title|body|string|true|none|
|» body|body|string|true|none|
|» category|body|string|true|none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": {
    "articleId": 1,
    "title": "打大原始干只",
    "body": "体及动任观再题老会是十提作还。律许七断角王光计世根回细该。研委基历每场光法过断各手代术和确。",
    "like": -4435938963185700
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|[%E6%96%87%E7%AB%A0%E4%BF%A1%E6%81%AF](#schema%e6%96%87%e7%ab%a0%e4%bf%a1%e6%81%af)|true|none|none|
|»» articleId|number|true|none|保证非负|
|»» title|string|true|none|标题|
|»» body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|»» like|number|true|none|保证非负|
|»» category|number|true|none|none|
|»» category_name|string|true|none|对应分类名|
|» msg|string|true|none|none|

## GET 搜索文章

GET /article/search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|category|query|string|false|不填默认不分类搜索~|
|title|query|string|false|none|
|body|query|string|false|none|
|limit|query|string|false|获取列表时的分页|
|offset|query|string|false|获取列表时的偏移|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": [
    27,
    70,
    57
  ],
  "msg": "Lorem ipsum"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» code|integer|true|none|none|
|» data|[[%E6%96%87%E7%AB%A0%E4%BF%A1%E6%81%AF](#schema%e6%96%87%e7%ab%a0%e4%bf%a1%e6%81%af)]|true|none|none|
|»» articleId|number|true|none|保证非负|
|»» title|string|true|none|标题|
|»» body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|»» like|number|true|none|保证非负|
|»» category|number|true|none|none|
|»» category_name|string|true|none|对应分类名|
|» msg|string|true|none|none|

## POST 点赞

POST /article/like

toggle模式，如果没点赞就给他点上，如果已经点赞了就给他取消了~

> Body 请求参数

```yaml
articleid: "0"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» articleid|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_用户信息">用户信息</h2>

<a id="schema用户信息"></a>
<a id="schema_用户信息"></a>
<a id="tocS用户信息"></a>
<a id="tocs用户信息"></a>

```json
{
  "name": "string",
  "status": "string",
  "userid": 0,
  "gender": "string",
  "avatar": "string",
  "desc": "string",
  "school": "string"
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|name|string|true|none|none|
|status|string|true|none|none|
|userid|integer|true|none|保证非负|
|gender|string|true|none|none|
|avatar|string|true|none|none|
|desc|string|true|none|none|
|school|string|true|none|none|

<h2 id="tocS_文章信息">文章信息</h2>

<a id="schema文章信息"></a>
<a id="schema_文章信息"></a>
<a id="tocS文章信息"></a>
<a id="tocs文章信息"></a>

```json
{
  "articleId": 0,
  "title": "string",
  "body": "string",
  "like": 0,
  "category": 0,
  "category_name": "string"
}

```

### 属性

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|articleId|number|true|none|保证非负|
|title|string|true|none|标题|
|body|string|true|none|正文，推荐前端使用markdown编辑器，后端不对正文部分进行异常校验|
|like|number|true|none|保证非负|
|category|number|true|none|none|
|category_name|string|true|none|对应分类名|

