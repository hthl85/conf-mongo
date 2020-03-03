package mongoconf

// MongoConf struct
type MongoConf struct {
	Host          string
	Username      string
	Password      string
	Dbname        string
	Port          uint64
	MinPoolSize   uint64
	MaxPoolSize   uint64
	MaxIdleTimeMS uint64
	Timeout       int
}
