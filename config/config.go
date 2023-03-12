package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Pgsql   Pgsql   `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Mongo   Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
}

type System struct {
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                               // 端口
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"img-height" yaml:"img-height"` // 验证码高度
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type Mongo struct {
	DB   string `mapstructure:"db" json:"db" yaml:"db"`       // mongodb数据库
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址:端口
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 密码
}
type Zap struct {
	Level         string   `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string   `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string   `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string   `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string   `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string   `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	OutputPaths   []string `mapstructure:"outputPaths" json:"outputPaths" yaml:"outputPaths"`
	MaxAge        int      `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine      bool     `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole  bool     `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

type GeneralDB struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

type Pgsql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 基于配置文件获取 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

func (m *Pgsql) GetLogMode() string {
	return m.LogMode
}
