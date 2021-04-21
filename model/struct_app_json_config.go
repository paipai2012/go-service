package model

type AppInfo struct {
	AppName    string `json:"app_name"`
	AppMode    string `json:"app_mode"`
	DriverName string `json:"driver_name"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	DataBase   string `json:"database"`
}
