package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/tgo-team/tgo-talkapi/config"
	"github.com/tgo-team/tgo-talkapi/utils"
	"time"
)


type DB struct {
	DbHost string
	DB string
	Username string
	Password string
	conn *dbr.Connection
}

var d *DB

func Init(config config.MysqlConfig)   {
	if d==nil {
		d = &DB{
			DbHost:config.Addr,
			DB:config.Db,
			Username:config.User,
			Password:config.Password,
		}
		d.conn = d.createConn()
	}
}

//初始化MYSQL
func (d *DB) createConn() *dbr.Connection {

	fmt.Println("init mysql...")
	loc,_ := time.LoadLocation("Local")
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",d.Username,d.Password,d.DbHost,d.DB,loc.String())
	fmt.Println(connInfo)
	var err error
	conn,err := dbr.Open("mysql",connInfo,nil)

	utils.CheckErr(err)
	conn.SetMaxOpenConns(2000)
	conn.SetMaxIdleConns(1000)
	conn.SetConnMaxLifetime(time.Second*60*60*4) //mysql 默认超时时间为 60*60*8=28800 SetConnMaxLifetime设置为小于数据库超时时间即可
	conn.Ping()

	fmt.Println("数据库连接成功！")
	return  conn
}


func  NewSession() *dbr.Session {

	return d.conn.NewSession(nil)
}


const (
	timeFormart = "2006-01-02 15:04:05"
)

type Time time.Time
type BaseModel struct {
	Id int64 `json:"id"`
	CreatedAt Time `json:"created_at"`
	UpdatedAt Time `json:"updated_at"`
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}