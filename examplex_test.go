package xjson

import (
	"log"
	"testing"
	"time"
)

func ExampleEmptyArrayAndEmptyObject () {
	response := struct {
		Books []string `json:"books"`
		Map map[string]string `json:"map"`
	}{}
	data, err := Marshal(response) ; if err != nil {panic(err)}
	log.Print(string(data)) // {"books":null, map:{}}
}
func TestExampleEmptyArrayAndEmptyObject(t *testing.T) {
	emptyObjectAndArrayNotNull = true
	defer func() {
		emptyObjectAndArrayNotNull = false
	}()
	ExampleEmptyArrayAndEmptyObject()
}
func ExampleIntFloatAutoConvertString () {
	request := struct {
		Page int `json:"page"`
		Price float64 `json:"price"`
	}{}
	err := Unmarshal([]byte(`{"page":"1", "price": "1.05"}`), &request) ; if err != nil {panic(err)}
	log.Printf("%+v", request) // {Page:1 Price:1.05}
}
func TestExampleIntFloatAutoConvertString(t *testing.T) {
	autoNumberConvertString = true
	defer func() {
		autoNumberConvertString = false
	}()
	ExampleIntFloatAutoConvertString()
}

func ExampleChinaTime() {
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
}
func TestExampleChinaTime(t *testing.T) {
	ExampleChinaTime()
}