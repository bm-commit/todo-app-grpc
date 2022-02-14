package settings

import (
	"log"

	"github.com/spf13/viper"
	"github.com/todo-app/internal/logger"
)

// Config contains the global configuration
type Config struct {
	GrpcAPIPort    uint   `mapstructure:"grpc-api-port" yaml:"grpc-api-port"`
	HTTPAPIPort    uint   `mapstructure:"http-api-port" yaml:"http-api-port"`
	HTTPMetricPort uint   `mapstructure:"http-metric-port" yaml:"http-metric-port"`
	JwtSecret      string `mapstructure:"jwt-secret" yaml:"jwt-secret"`
	Logger         Logger `mapstructure:"logger" yaml:"logger"`
	DBconn         DBConn `mapstructure:"dbconn"`
}

// Logger contains the log configuration.
type Logger struct {
	JSON        bool   `mapstructure:"json" yaml:"json"`
	FileEnabled bool   `mapstructure:"file_enable" yaml:"file_enable"`
	Level       string `mapstructure:"level" yaml:"level"`
	FilePath    string `mapstructure:"file_path" yaml:"file_path"`
}

// SetDefaults sets the default values.
func (l *Logger) SetDefaults() {
	l.JSON = true
	l.FileEnabled = false
	l.Level = "info"
	l.FilePath = "logs/todoapp.log"
}

// DBConn contains the Database configuration.
type DBConn struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Database string `mapstructure:"database"`
	Schema   string `mapstructure:"schema"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

// SetDefault sets the default values.
func (db *DBConn) SetDefault() {
	db.Host = "localhost"
	db.Port = "1434"
	db.User = "dev"
	db.Pass = "dev"
	db.Database = "Todoapp"
	db.Schema = "dbo"
	db.SSLMode = false
}

func init() {
	// Global
	viper.SetDefault("grpc-api-port", 51001)
	viper.SetDefault("http-api-port", 3000)
	viper.SetDefault("http-metric-port", 2112)
	viper.SetDefault("jwt-secret", "MySecret")

	// DBConn
	var dbconn DBConn
	dbconn.SetDefault()
	viper.SetDefault("dbconn", dbconn)

	// Logger
	var l Logger
	l.SetDefaults()
	viper.SetDefault("logger", l)
}

// LoadConfiguration loads the configuration from config file.
// If the load configuration fails, it loads the default configuration
// and creates a config file with the defaults values.
func LoadConfiguration(path string) *Config {
	logger.Info("loading configuration...")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("cannot load the configuration %v, creating default configuration", err)
		if err = viper.SafeWriteConfigAs(path); err != nil {
			log.Fatalf("failed to create default configuration.")
		}
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("failed unmarshall configuration: %v", err)
	}

	logger.NewLogger(logger.Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: config.Logger.JSON,
		ConsoleLevel:      logger.GetLevel(config.Logger.Level),
		EnableFile:        config.Logger.FileEnabled,
		FileJSONFormat:    config.Logger.JSON,
		FileLevel:         logger.GetLevel(config.Logger.Level),
		FileLocation:      config.Logger.FilePath,
	})

	return config
}
