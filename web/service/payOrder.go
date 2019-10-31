package service
import(
	"longchain.com/memoriae/profit/config"
	db "longchain.com/memoriae/profit/db/mongodb"
)
var(
	dbPayOrder = db.New(config.MongodbUrl,"test","payOrder")
)
type PayOrder struct{
	// Tx string
	// Address string
	Timestamp int `json:"timestamp"`
	// Type string
	// Role string
	// From string
	// To string
	Value string `json:"value"`
}
// 按角色查询收益
func Profit(role,addr string) []PayOrder {
	res := []PayOrder{}
	dbPayOrder.Query(db.M{"role":role,"to":addr},&res)
	return res
}