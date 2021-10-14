package settings

var defaultConfig = map[string]interface{}{
	"server.port":       8080,
	"db.dataSourceName": "root:root.com@tcp(127.0.0.1:3306)/lufflysex?charset=utf8&parseTime=True&loc=Local",
}
