package mongodb
import(
	"testing"
)
var (
	db = New("127.0.0.1:27017","eth","block")
	stu1 = Student{
		Name:  "zhangsan",
		Phone: "13480989765",
		Email: "329832984@qq.com",
		Sex:   "F1",
	}
)
type Student struct {
	// _Id bson.ObjectId `bson:"_id"`
	Name string
	Phone string
	Email string
	Sex string
}

func Test1(t *testing.T) {
	db.Insert(stu1)
}
func Test2(t *testing.T) {
	res := []Student{}
	db.Query(M{"name": "zhangsan"},&res)
	t.Log(res)
}
func Test3(t *testing.T) {
	db.Update(M{"phone": "13480989765"},M{"phone":"666"})
}
