package xjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init () {
	emptyObjectAndArrayNotNull = false
	autoNumberConvertString = false
}


func TestNilSliceConvertEmptyArray(t *testing.T) {
	emptyObjectAndArrayNotNull = true
	defer func() {
		emptyObjectAndArrayNotNull = false
	}()
	a := struct {
		Bools []bool
		Ints []int
		Structs []struct{A int}
		Bytes []byte
		Strings []string
		Map map[int]int
	}{}
	data, err := Marshal(a) ; if err != nil {panic(err)}
	actual := string(data)
	expect := `{"Bools":[],"Ints":[],"Structs":[],"Bytes":[],"Strings":[],"Map":{}}`
	assert.Equal(t,expect, actual)
}



func TestIntConvertString (t *testing.T) {
	autoNumberConvertString = true
	defer func() {
		autoNumberConvertString = false
	}()
	query := struct {
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Float32 float32
		Float64 float64
	}{}
	err := Unmarshal([]byte(`{"Int":"1","Int8":"2","Int16":"3","Int32":"4","Int64":"5","Uint":"1","Uint8":"2","Uint16":"3","Uint32":"4","Uint64":"5","Float32": "1.111", "Float64": "2.222"}`),&query)
	assert.NoError(t, err)
	assert.Equal(t,1, query.Int)
	assert.Equal(t,int8(2), query.Int8)
	assert.Equal(t,int16(3), query.Int16)
	assert.Equal(t,int32(4), query.Int32)
	assert.Equal(t,int64(5), query.Int64)
	assert.Equal(t,uint(1), query.Uint)
	assert.Equal(t,uint8(2), query.Uint8)
	assert.Equal(t,uint16(3), query.Uint16)
	assert.Equal(t,uint32(4), query.Uint32)
	assert.Equal(t,uint64(5), query.Uint64)
	assert.Equal(t,float32(1.111), query.Float32)
	assert.Equal(t,float64(2.222), query.Float64)

	{
		a := struct {
			A int
		}{}
		err := Unmarshal([]byte(`{"A":"xx"}`), &a)
		assert.EqualError(t, err, "json: cannot unmarshal string into Go struct field .A of type int")
	}
}

func TestChinaTime(t *testing.T) {
	String := func (v interface{}) string {
		data, err := Marshal(v) ; if err != nil {panic(err)}
		return string(data)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 2*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 21:29:29"`)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 8*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 15:29:29"`)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 0*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 23:29:29"`)
	}
	{
		type Request struct {
			Time ChinaTime `db:"time"`
		}
		req := Request{}
		err := Unmarshal([]byte(`{"time":"2020-07-31 15:37:44"}`), &req)
		assert.NoError(t, err)
		assert.Equal(t, req.Time.In(time.FixedZone("CST", 8*3600)).String(), "2020-07-31 15:37:44 +0800 CST")
	}
	{
		type Reply struct {
			Time ChinaTime `db:"time"`
		}
		reply := Reply{}
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.UTC)
		assert.Equal(t, err, nil)
		reply.Time = NewChinaTime(tValue)
		assert.Equal(t, reply.Time.String(), "2020-07-31 23:29:29 +0800 CST")
	}
}