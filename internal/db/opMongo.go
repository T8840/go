package db

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

 	 mgo "godemo/lib/db/mongodb"
)

// ------------------- Level Struct -----------------------
type Level struct {
	Id      int64    `json:"id"`
	Cnt     int64    `json:"cnt"`
	Course  []string `json:"course"`
	Teacher []string `json:"teacher"`
}

// GetList 获取全量的数据
func (l *Level) GetList(filter interface{}, objs interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return err
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	err = cur.All(context.Background(), &objs)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", objs)

	return nil
}

// AddOne 新增一条数据
func (l *Level) AddOne() (*mongo.InsertOneResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return nil, err
	}

	objId, err := collection.InsertOne(context.TODO(), *l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("录入数据成功，objId:", objId)

	return objId, nil
}

// EditOne 编辑一条数据
func (l *Level) EditOne(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return err
	}

	update := bson.M{"$set": l}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println("collection.UpdateOne:", updateResult)

	return nil
}

// 更新数据 - 存在更新，不存在就新增
func (l *Level) Update(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return err
	}

	update := bson.M{"$set": l}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("collection.UpdateOne:", updateResult)

	return nil
}

// 删除一条匹配的数据
func (l *Level) Del(filter interface{}) (*mongo.DeleteResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return nil, err
	}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("collection.DeleteOne:", deleteResult)

	return deleteResult, nil
}

// 删除全部匹配的数据
func (l *Level) DelAll(filter interface{}) (*mongo.DeleteResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return nil, err
	}

	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("collection.DeleteOne:", deleteResult)

	return deleteResult, nil
}

// 删除整个文档
func (l *Level) DelTable() error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return nil
	}

	err = collection.Drop(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Sectle 模糊查询
// bson.M{"name": primitive.Regex{Pattern: "深入"}}
func (l *Level) Sectle(filter interface{}, objs interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return err
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	ss := []Level{}
	for cur.Next(context.Background()) {
		var s Level
		if err = cur.Decode(&s); err != nil {
			fmt.Println(err)
		}
		//log.Println("collection.Find name=primitive.Regex{xx}: ", s)
		ss = append(ss, s)
	}
	cur.Close(context.Background())
	objs = ss

	return nil
}

// 统计collection的数据总数
func (l *Level) Count(filter interface{}) (int, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return 0, err
	}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Println(count)
	}
	log.Println("collection.CountDocuments:", count)

	return int(count), nil
}

// 准确搜索一条数据
func (l *Level) GetOne(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_CLASS_NAME)
	if err != nil {
		return err
	}

	err = collection.FindOne(context.Background(), filter).Decode(l)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println("collection.FindOne: ", l)
	return nil
}



// ------------------- Student Struct -----------------------


type Student struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

// GetList 获取全量的数据
func (s *Student) GetList(filter interface{}, objs interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return err
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	err = cur.All(context.Background(), &objs)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", objs)

	return nil
}

// AddOne 新增一条数据
func (s *Student) AddOne() (*mongo.InsertOneResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return nil, err
	}

	objId, err := collection.InsertOne(context.TODO(), *s)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("录入数据成功，objId:", objId)

	return objId, nil
}

// EditOne 编辑一条数据
func (s *Student) EditOne(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return err
	}

	update := bson.M{"$set": s}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println("collection.UpdateOne:", updateResult)

	return nil
}

// 更新数据 - 存在更新，不存在就新增
func (s *Student) Update(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return err
	}

	update := bson.M{"$set": s}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("collection.UpdateOne:", updateResult)

	return nil
}

// 删除一条匹配的数据
func (s *Student) Del(filter interface{}) (*mongo.DeleteResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return nil, err
	}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("collection.DeleteOne:", deleteResult)

	return deleteResult, nil
}

// 删除全部匹配的数据
func (s *Student) DelAll(filter interface{}) (*mongo.DeleteResult, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return nil, err
	}

	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("collection.DeleteOne:", deleteResult)

	return deleteResult, nil
}

// 删除整个文档
func (s *Student) DelTable() error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return nil
	}

	err = collection.Drop(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Sectle 模糊查询
// bson.M{"name": primitive.Regex{Pattern: "深入"}}
func (s *Student) Sectle(filter interface{}, objs interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return err
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	ss := []Student{}
	for cur.Next(context.Background()) {
		var s Student
		if err = cur.Decode(&s); err != nil {
			fmt.Println(err)
		}
		//log.Println("collection.Find name=primitive.Regex{xx}: ", s)
		ss = append(ss, s)
	}
	cur.Close(context.Background())
	objs = ss

	return nil
}

// 统计collection的数据总数
func (s *Student) Count(filter interface{}) (int, error) {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return 0, err
	}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Println(count)
	}
	log.Println("collection.CountDocuments:", count)

	return int(count), nil
}

// 准确搜索一条数据
func (s *Student) GetOne(filter interface{}) error {
	collection, err := mgo.Get_collecton(mgo.TBL_STUDENT_NAME)
	if err != nil {
		return err
	}

	err = collection.FindOne(context.Background(), filter).Decode(s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println("collection.FindOne: ", s)
	return nil
}

// ------------------- Level Opration -----------------------

func Insert() {
	lev1 := &Level{Id: 100, Cnt: 6, Teacher: []string{"王老师", "张老师", "李老师", "孙老师", "赵老师"}}
	lev1.AddOne()
	// lev1.GetList()
	fmt.Println("\r\n===============读取班级信息=================")
	ll := &Level{}
	filter00 := bson.M{"id": 100}
	levels := []Level{}
	err := ll.GetList(filter00, levels)
	if err == nil {
		for _, l := range levels {
			fmt.Println(l)
		}
	}
}

// ------------------- Student Opration -----------------------

func InsertToStudent() {
	var pt mgo.Table_interface

	pt = &Student{Id: 101291892}
	pt.AddOne()

	fmt.Println("===============插入2条学生信息=================")
	stu := &Student{Id: 100}
	stu1 := &Student{Id: 10, Name: "晴天公子", Age: 12}
	stu2 := &Student{Id: 10, Name: "沫沫", Age: 12, Height: 168, Weight: 60}
	stu3 := &Student{Id: 10, Name: "黛玉晴雯子", Age: 12}

	pt = stu
	pt.AddOne()
	stu1.AddOne()
	stu2.AddOne()
	stu3.AddOne()

	fmt.Println("\r\n===============读取学生信息=================")
	filter0 := bson.M{"name": "王葛"}
	students := []Student{}
	err := pt.GetList(filter0, students)
	if err == nil {
		for _, v := range students {
			fmt.Println(v)
		}
	}

	fmt.Println("\r\n=============模糊查找========================")
	filter1 := bson.M{"name": primitive.Regex{Pattern: "晴"}}
	err = pt.Sectle(filter1, students)
	if err == nil {
		for _, s := range students {
			fmt.Println(s)
		}
	}

	// fmt.Println("\r\n===========删除1条学生信息===================")
	// filter2 := bson.M{"name": "王葛", "age": 12}
	// res, err := pt.Del(filter2)
	// if err == nil {
	// 	fmt.Printf("删除数据条数为：%d条，删除的数据为：%+v", res.DeletedCount, filter2)
	// }

	// filter3 := bson.M{"age": 12}
	// err = pt.GetList(filter3, students)
	// if err == nil {
	// 	for _, v := range students {
	// 		fmt.Println(v)
	// 	}
	// }

	// fmt.Println("\r\n==================编辑一条数据===================")
	// filter4 := bson.M{"age": 12}
	// stu2.Age = 22110
	// stu2.EditOne(filter4)

	// fmt.Println("\r\n==================统计数据===================")
	// filter5 := bson.M{"age": 22110}
	// sstu := &Student{}
	// cnt, err := sstu.Count(filter5)
	// if err == nil {
	// 	fmt.Printf("filter5=%+v, cnt=%+v\r\n", filter5, cnt)
	// }

	// fmt.Println("\r\n==================删除学生整表===================")
	//pt.DelTable()

	//pt = lev1
	//pt.DelTable()
}      

		

	


