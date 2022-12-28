package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"
	"sync"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	// conf "godemo/utils/conf"
	// logger "godemo/utils/logger"

)

type Table_interface interface {
	GetList(filter interface{}, objs interface{}) error
	GetOne(filter interface{}) error
	AddOne() (*mongo.InsertOneResult, error)
	EditOne(filter interface{}) error
	Update(filter interface{}) error
	Del(filter interface{}) (*mongo.DeleteResult,error)
	DelAll(filter interface{}) (*mongo.DeleteResult,error)
	DelTable() error
	Sectle(filter interface{},objs interface{}) error
	Count(filter interface{}) (int,error)
}

type objs interface{}

var (
	DB_URI                   = "mongodb://admin:admin@127.0.0.1:27017" //  带账号名的链接
	DB_NAME                  = "admin"                                     // 数据库名
	maxTime                  = time.Duration(5)                           // 链接超时时间
	DB_CLIENT_MAX_NUM uint64 = 100                                        // 链接数
	db_mutex          sync.Mutex
	db_ready          = false
	toDB              *mongo.Database              // database 话柄
	collection_map    map[string]*mongo.Collection // collection 话柄
)

// TODO: 调整为配置
var (
	TBL_STUDENT_NAME         = "student"                                  // 学生表名
	TBL_CLASS_NAME           = "class"                                    // 班级表名
	TBL_SITE_NAME            = "safty_site"
)


// pool 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	opts := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	opts.SetMaxPoolSize(num)
	// 发起链接
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}

func init() {
	var err error
	collection_map = make(map[string]*mongo.Collection)

	condb := func() error {
		db_mutex.Lock()
		defer db_mutex.Unlock()
		fmt.Println("----------------------ready connect to mongodb!")
		toDB, err = ConnectToDB(DB_URI, DB_NAME, maxTime, DB_CLIENT_MAX_NUM)
		if err != nil {
			db_ready = false
			fmt.Println("----------------------failed connectto mongodb!")
			return err
		}
		fmt.Println("----------------------connectted to mongodb!")

		db_ready = true

		// TODO: 调整为配置
		collection_map[TBL_STUDENT_NAME] = toDB.Collection(TBL_STUDENT_NAME)
		collection_map[TBL_CLASS_NAME] = toDB.Collection(TBL_CLASS_NAME)
		collection_map[TBL_SITE_NAME] = toDB.Collection(TBL_SITE_NAME)

		return nil
	}
	condb()

	go func() {
		for {
			timeout := time.After(time.Second * 4)
			<-timeout
			if toDB != nil {
				if err := toDB.Client().Ping(context.Background(), readpref.Primary()); err != nil {
					fmt.Println(err)

					fmt.Println(" retry connect db")
					condb()
				}
			} else {
				condb()
			}
		}
	}()
}

func Exit_DB() {
	if toDB != nil {
		toDB.Client().Disconnect(context.Background())
	}
}

func Get_collecton(TableName string) (*mongo.Collection, error) {
	db_mutex.Lock()
	defer db_mutex.Unlock()

	if !db_ready {
		return nil, errors.New("db disconnected, waitting for retry connect")
	}

	collection, ok := collection_map[TableName]
	if !ok {
		fmt.Println("can't find collection")
		return nil, errors.New("can't find collection")
	}

	return collection, nil
}

func BuildFilter(ftype string, filterData interface{}) (interface{}, error) {
	filter, err := func(str string) (interface{}, error) {
		switch str {
		case "D":
			return bson.D{}, nil
		case "E":
			return bson.E{}, nil
		case "M":
			return bson.M{}, nil
		case "A":
			return bson.A{}, nil
		}

		return nil, errors.New("unknow ftype, when BuildQueryFilter()")
	}(ftype)
	if err != nil {
		return nil, err
	}

	data, err := bson.Marshal(filterData)
	if err != nil {
		fmt.Printf("marshal error: %v", err)
		return filter, err
	}

	err = bson.Unmarshal([]byte(data), filter)
	if err != nil {
		fmt.Printf("unmarshal error: %v", err)
		return filter, err
	}
	log.Printf("filter: %v", filter)
	return filter, nil
}

