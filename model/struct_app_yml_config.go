package model

type AppConfig struct {
	App struct {
		Name string `yaml:"name" json:"name"`
		Mode string `yaml:"mode" json:"mode"`
	}

	MySql struct {
		Host       string `yaml:"host" json:"host"`
		Port       string `yaml:"port" json:"port"`
		DriverName string `yaml:"driverName" json:"driverName"`
		UserName   string `yaml:"username" json:"username"`
		Password   string `yaml:"password" json:"password"`
		DataBase   string `yaml:"database" json:"database"`
	}

	Redis struct {
		Host string `yaml:"host" json:"host"`
	}
}
