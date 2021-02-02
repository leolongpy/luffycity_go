package main

//解析ini文件
type server struct {
	IP   string `ini:"ip"`
	Port int64  `ini:"port"`
}
type mysql struct {
	Host     string `ini:"host"`
	Port     int64  `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type config struct {
	Server server `ini:"server"`
	MySql  mysql  `ini:"mysql"`
}

// Load 加载配置文件

func main() {

}
