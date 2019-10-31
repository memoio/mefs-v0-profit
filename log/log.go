package log
import(
	"runtime"
	"fmt"
	"time"
	"strings"
	"os"
)
/*
可以打印调用路径
*/
func date()(string){
  timeLayout := "2006-01-02 15:04:05"
  date := time.Now().Format(timeLayout)
  return date
}
// 包名/文件名
func formatFileName(str string) string {
	res := ""
	i := strings.Index(str,"/src/")
	if i != -1 {
		res = str[i+5:]
	}
	return res
}
// 0000-00-00 00:00:00 级别 说明 文件名 行
func out(level string, msg ...interface{}){
	skip := 2
	_, fileName, line, ok := runtime.Caller(skip)
	if ok {
		fmt.Printf("%s %s %v %s %d\n",date(),level,msg,formatFileName(fileName),line)
	}else{
		fmt.Printf("%s %s %v %s %d\n",date(),level,msg,"",0)
	}
}
func Debug(msg ...interface{}){
	out("debug:",msg...)
}
func Info(msg ...interface{}){
	out(" info:",msg...)
}
func Warn(msg ...interface{}){
	out(" warn:",msg...)
}
func Error(msg ...interface{}){
	out("error:",msg...)
}
func Fatal(msg ...interface{}){
	out("fatal:",msg...)
	os.Exit(1)
}
func Panic(msg ...interface{}){
	out("panic:",msg...)
	panic(msg)
}