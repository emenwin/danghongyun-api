# danghongyun-api-golang

当虹云api ，golang 非官方版本


* 点播API https://developer.danghongyun.com/rest_api_new.html
* 直播API https://developer.danghongyun.com/rest_api_live.html
* UGC直播 https://developer.danghongyun.com/rest_api_ugc.html


## 参数说明

虹视云的API，一般都会包含5个参数action，accessKey，version，timestamp 和 signature

```
action：api名称

accessKey：分配给用户的访问key。

version：api版本

timestamp：时间戳(1970年1月1日以来的毫秒数，格林尼治时间)

signature：请求签名，签名生成规则：对所有参数按照名称排序(CASE_INSENSITIVE_ORDER)生成一个字符串，并在字符串头上加上accessSecret，然后对该字符串调用[HmacSHA256](http://en.wikipedia.org/wiki/Hash-based_message_authentication_code) 算法生成signature。每一个API，根据输入参数的不同，生成的签名都是不同的。

```

我们以API：创建频道 为例来说明签名的生成，该接口的参数为accessKey，action，signature，version，timestamp，name:

```
accessKey = "a020e193-0f1"
accessSecret = "5GcXHNYdAVVdFW0yervG"
action = "ugcCreateChannel"
version = "2.0"
timestamp = "1466488681033"
name = "直播测试频道"

```

那么调用该接口的signature参数就是：

```
 hmacSHA256Encrypt("5GcXHNYdAVVdFW0yervGaccessKey=a020e193-0f1action=ugcCreateChannelname=直播测试频道timestamp=1466488681033version=2.0", 5GcXHNYdAVVdFW0yervG)
```

计算后的值为

```
69a22291e517179e76e55ad7c21269c17462e102adf0eb0526cc5012d9b4cd9a
```

最终的请求为:

```
http://api.danghongyun.com/rest?timestamp=1466488681033&accessKey=a020e193-0f1&name=直播测试频道&action=ugcCreateChannel&signature=69a22291e517179e76e55ad7c21269c17462e102adf0eb0526cc5012d9b4cd9a&version=2.0

```

## Demo

