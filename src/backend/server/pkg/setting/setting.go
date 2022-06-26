package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

type ServerSettings struct {
	RunMode  string
	HttpPort string
}

type AppSettings struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
}

type DatabaseSettings struct {
	DBType     string
	UserName   string
	Password   string
	Host       string
	DBName     string
	Charset    string
	ParserTime bool
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSettings(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}

	return nil
}
