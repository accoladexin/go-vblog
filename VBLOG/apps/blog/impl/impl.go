package impl

import (
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/conf"
	"github.com/accoladexin/vblog/ioc"
	"gorm.io/gorm"
)

// 先定义一个对象, 由这个对象来实现业务功能
// impl 才实现了业务接口生成
// 现在 impl对象 充当的是 mvc里面的控制器
// 控制器实现了一个核心业务方法:  CreateBlog
// 依赖myql的配置, 怎么访问MySQL
type impl struct {
	// ORM对象
	db *gorm.DB
}

// 现在这个业务实现的包, 封装完整, 并且不对外开发, 只能通过ioc访问
// 保证了业务实现的安全
// 这个对象满足2从身份: 业务服务/Ioc Controller托管

// 声明一个变了_, 为啥声明的一个变量又不用他的指，只想要他的接口约束
var _ blog.Service = &impl{}
var _ ioc.ServiceObjIoc = &impl{}

// 业务实例如何访问配置
// 通常为这个实例类 提供初始化方法
// 在配置初始化好后, 再调用
func (i *impl) Init() error {
	// conf.C().MySQL 读取config的mysql配置，调用初始化数据库方方法
	// ORM() 获取全局的ORM对象
	orm := conf.C().MySQL.ORM()
	// 赋值
	i.db = orm.Debug()
	return nil
}

func NewImpl() *impl {
	return &impl{}
}
func (i *impl) Name() string {
	return blog.AppName
}
func init() {
	ioc.RegistryServiceIoc(&impl{})
}
