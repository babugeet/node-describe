package constants

var cfgFile string

func SetCfgFile(file string) {
	cfgFile = file
}

func GetCfgFile() string {
	return cfgFile
}
