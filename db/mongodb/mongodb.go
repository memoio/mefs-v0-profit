package mongodb
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"longchain.com/memoriae/profit/log"
)

type Mongo struct {
	db *mgo.Collection
}
type M bson.M
// url 连接地址 127.0.0.1:27017
// dbName 数据库名称
// tbName 表名
func New(url,dbName,tbName string) Mongo {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(tbName)
	mongo := Mongo{
		db:c,
	}
	return mongo
}
// 插入一条或多条记录
func (this Mongo) Insert(data ...interface{}) {
	err := this.db.Insert(data...)
	if err != nil {
		log.Error(err)
	}
}
func (this Mongo) Query(where M, res interface{}) {
	err := this.db.Find(where).All(res)
	if err != nil {
		log.Error(err)
	}
}
func (this Mongo) QueryOne(where M, res interface{}) {
	err := this.db.Find(where).One(res)
	if err != nil {
		log.Error(err)
	}
}
func (this Mongo) Update(where M, data interface{}) {
	_,err := this.db.UpdateAll(where, M{"$set": data})
	if err != nil {
		log.Error(err)
	}
}
func (this Mongo) Delete(where M) {
	_, err := this.db.RemoveAll(where)
	if err != nil {
		log.Error(err)
	}
}