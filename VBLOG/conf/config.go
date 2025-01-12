package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	orm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 默认的配置
//func DefaultConfig() *Config {
//	return &Config{
//		MySQL: newDefaulMySQL(),
//
//	}
//}

type Config struct {
	MySQL *MySQL `json:"mysql" toml:"mysql"` // 对应配置文件的[mysql]
	Http  *Http  `json:"http" toml:"http"`
}

// 默认Config 配置
func DefaultConfig() *Config {
	return &Config{
		// 读取默认mysqlconfig
		MySQL: newDefaulMySQL(),
		Http:  newDefaultHttpConfig(),
	}

}

// config 类型的自定义fmt.print
// 格式化成一个json
func (c *Config) String() string {
	d, _ := json.MarshalIndent(c, "", "  ")
	return string(d)
}

// mysql 配置获取
type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`             // host = "127.0.0.1"
	Port     int    `json:"port" toml:"port"  env:"MYSQL_Port`             // port = 3306
	DB       string `json:"db" toml:"db"  env:"MYSQL_DB`                   // db = "test"
	Username string `json:"username" toml:"username"  env:"MYSQL_USERNAME` ////username = "root"
	Password string `json:"password" toml:"password"  env:"MYSQL_PASSWORD` ////password = "123456"

	// 高级参数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	// 面临并发安全
	lock sync.Mutex
	// 数据库连接
	db *gorm.DB
}

// 默认Mysql的配置
func newDefaulMySQL() *MySQL {
	return &MySQL{
		Host:     "47.109.189.135",
		Port:     3306,
		DB:       "vblog",
		Username: "root",
		Password: "Gzdx123456@",
	}
}

// 获取数据库连接池
func (m *MySQL) GetConnPool() (*sql.DB, error) {
	var err error
	// multiStatements 让db 可以执行多个语句 select; insert;
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&multiStatements=true&parseTime=True&loc=Local&multiStatements=true",
		m.Username, m.Password, m.Host, m.Port, m.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Connect mysql <%s> error: %s", dsn, err)
	}
	// 对连接池设计
	if m.MaxOpenConn != 0 {
		db.SetMaxOpenConns(m.MaxOpenConn)
	}
	if m.MaxIdleConn != 0 {
		db.SetMaxIdleConns(m.MaxIdleConn)
	}
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}

	return db, nil
}

// ORM 获取orm对象
// 先获取一个原生的mysql初始化对象
// 然后初始化orm对象
// 也可以参照https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func (m *MySQL) ORM() *gorm.DB {

	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		// 初始化db
		pool, err := m.GetConnPool()
		if err != nil {
			panic(err)
		}
		// 1.2 使用pool 初始化orm db对象
		m.db, err = gorm.Open(
			// 初始化DB
			// 1.1 获取sql.DB
			orm_mysql.New(orm_mysql.Config{
				Conn: pool,
			}),
			&gorm.Config{
				// 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
				PrepareStmt: true,
				// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。
				// 但这会降低性能，如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升
				SkipDefaultTransaction: true,
				// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法
				// CreateBatchSize: 200,
			})
		if err != nil {
			panic(err)
		}
	}
	return m.db

}
