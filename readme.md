# jsonx

```go
import "github.com/og/jsonx"
```
Go JSON standard package extension library.

> Go 官方 JSON 标准库扩展

## Extension feature

> 扩展特性

1. nil slice nad nil map marshal to `[]` or `{}` instead of null.
> 1. 空切片空 map 序列化字符串为 `[]` 或 `{}` 而不是 `null`。


```go
response := struct {
    Books []string `json:"books"`
    Map map[string]string `json:"map"`
}{}
data, err := jsonx.Marshal(response) ; if err != nil {panic(err)}
log.Print(string(data)) // {"books":null, map:{}}
```

2. auto int and float convert to stirng instead return error.

> 自动int和 float到字符串，而不是返回错误

```go
request := struct {
    Page int `json:"page"`
    Price float64 `json:"price"`
}{}
err := jsonx.Unmarshal([]byte(`{"page":"1", "price": "1.05"}`), &request) ; if err != nil {panic(err)}
log.Printf("%+v", request) // {Page:1 Price:1.05}
```

3. ChinaTime
 
🇨🇳

支持将 `2020-12-10 15:45:35` 格式的时间以中国时区解析为 `time.Time`

```go
{
    request := struct {
        SendTime ChinaTime `json:"sendTime"`
    }{}
    err := Unmarshal([]byte(`{"sendTime": "2020-12-10 15:39:25"}`), &request) ; if err != nil {panic(err)}
    log.Print(request.SendTime.String()) // 2020-12-10 15:39:25 +0800 CST
}
{
    response := struct {
        CreateTime ChinaTime `json:"createTime"`
    }{
        CreateTime: NewChinaTime(time.Date(2000,1,1,0,0,0,0, time.UTC)),
    }
    data, err := Marshal(response) ; if err != nil {panic(err)}
    log.Print(string(data)) // {"createTime":"2000-01-01 08:00:00"}
}
```