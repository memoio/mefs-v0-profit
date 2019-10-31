package json
import(
	"strings"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"longchain.com/memoriae/profit/log"
)
type Json struct {
	data []byte
}

func NewByStr(str string) Json{
	return Json{
		data:[]byte(str),
	}
}
func (j Json) GetString(selector string) string{
	return jsoniter.Get(j.data,format(selector)...).ToString()
}
// 获取数组的长度
func (j Json) GetLength(selector string) int{
	res := jsoniter.Get(j.data,format(selector)...).ToString()

	return len(ToArray(res))
}


func split(data string) []string {
  res := ""
  c := ""
  
  for i:=0; i<len(data); i++ {
    c = string(data[i])
    if c == "." || c == "[" {
      res += ","
    }else if c == "]" {
      res += ""
    }else{
      res += c
    }
  }
  
  return strings.Split(res,",")
}
// 把 result.logs[0].topics[0] 格式化成 ["result","logs",0,"topics",0]
func format(data string) []interface{} {
	arr := split(data)
	res := make([]interface{}, len(arr))
	// 把string数组转成interface{}数组
	for i:=0; i<len(arr); i++ {
		// 如果可以转成数字则
		if val, err := strconv.Atoi(arr[i]); err == nil {
			res[i] = val
		}else{
			res[i] = arr[i]
		}
	}
	return res
}

// 压缩json字符串,为一行
// 引号中的空格不动
// 其它空格,换行删除
func Compress(data string) string{
  // true:当前在引号中
	flag := false
  res := ""
  c := ""
  for i:=0; i<len(data); i++ {
    c = string(data[i])
    if c == `"` {
      if flag {
        flag = false
      }else{
        flag = true
      }
    }
    if !flag {
      if !(c == " " || c == "\n" || c == "\t") {
        res += c
      }
    }else{
      res += c
    }
  }
  return res
}
// 把非标准的json转化成标准的
// 主要是给key加引号
func ToStandard(data string) string{
	data = Compress(data)
  // 如果是 { 下一个不是 " 则添加
  // 如果是 , 下一个不是 " 则添加,排除[]
  res := ""
  c := ""
  // false:当前不在[]中
  flag := false
  // 标记是否添加 右" true:添加
  rightQuotes := false
  for i:=0; i<len(data)-1; i++ {
    c = string(data[i])
    res += c
    if c == `[` {
      flag = true
    }else if c == `]` {
      flag = false
    }else if (c == `{` || (c == `,` && !flag)) && string(data[i+1]) != `"` {
      res += `"`
      rightQuotes = true
    // 下一个是 : 则添加 右"
    }else if rightQuotes && string(data[i+1]) == `:`{
      res += `"`
    }
  }
  // 添加最后一个
  res += string(data[len(data)-1])
  
  return res
}
// json字符串转map
func ToMap(data string) map[string]interface{}{
	var res map[string]interface{}
	err := jsoniter.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Error("json转map失败", err)
	}
	return res
}
// json字符串转数组
func ToArray(data string) []interface{}{
	var res []interface{}
	err := jsoniter.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Error("json转数组失败", err)
	}
	return res
}
func ToJson(data interface{}) string{
	res, err := jsoniter.Marshal(data)
	if err != nil {
		log.Error("转json字符串失败")
		return ""
	}
	return string(res)
}
func ToInterface(data string) interface{}{
	var res interface{}
	err := jsoniter.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Error("json转接口失败", err)
	}
	return res
}
func ToObject(data string, v interface{}){
	err := jsoniter.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Error("json转对象失败", err)
	}
}