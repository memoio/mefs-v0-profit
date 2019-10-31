package resultModel

import (
	"testing"
)

func Test1(t *testing.T) {
	t.Log(Ok("成功"))
	t.Log(Error("失败"))
}