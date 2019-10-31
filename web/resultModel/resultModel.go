package resultModel

type ResultModel struct{
	Success bool
	Code int
	Msg string
	Data interface{}
}

func Ok(data interface{}) map[string]interface{}{
	return ResultModel{true,200,"",data}.toMap()
}
func Error(data interface{}) map[string]interface{}{
	return ResultModel{false,444,"",data}.toMap()
}

func (rm ResultModel) toMap() map[string]interface{}{
	return map[string]interface{}{
		"success":rm.Success,
		"code":rm.Code,
		"msg":rm.Msg,
		"data":rm.Data,
	}
}