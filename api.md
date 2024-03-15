# rankbp

## 请求

### 当请求成功时，返回以下内容

+ Response 200

  {
    "success": true,
    "data": {
      ...
    }
  }


### 当请求失败时，返回以下内容

+ Response 200

  {
    "success": false,
    "error": "报错信息"
  }



## 外部链接爬取

#### 获取英雄（opgg）https://www.op.gg/champions

单独python爬虫



### 英雄池部分

#### 更新英雄数据基础数据`POST: {base_url}/api/herosdata/base`

request

```
{
	"heros":[
		{
			"heroId":123,
			"name":"abc",
			"position1":"mid",
			"position2":"",//可留空
		}，
		{
		...
		}
		...
	]
}
```

response

```
{
	"success":true
}
```



#### 获取英雄池`GET: {base_url}/api/herospool/get`

request

```
无请求体
```

response

```
{
	"success":true,
	"data":[
		{
			"best":["a","b",...],
			"good":["a","b",...]
    	},
		{
			...
		},
	...
	]
}
```

#### 修改英雄池`POST: {base_url}/api/herospool/change`

request（heropool为长度为5的数组，从上到下依次为top，jug，mid，adc，sup）

```
{
	"heropool":[
		{
			"best":["a","b",...],
			"good":["a","b",...]
    	},
		{
			...
		},
	...
	]
}
```

response

```
{
	"success":true
}
```

