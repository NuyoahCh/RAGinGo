package config

import (
	"github.com/spf13/viper"
)

// Config 存储所有应用的配置
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	OpenAI   OpenAIConfig   `mapstructure:"openai"`
	VectorDB VectorDBConfig `mapstructure:"vectordb"`
	Storage  StorageConfig  `mapstructure:"storage"`
}

// ServerConfig 包含 HTTP 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

// OpenAIConfig 包含 OpenAI 相关配置
type OpenAIConfig struct {
	APIKey              string  `mapstructure:"api_key"`
	EmbeddingModel      string  `mapstructure:"embedding_model"`
	CompletionModel     string  `mapstructure:"completion_model"`
	MaxCompletionTokens int     `mapstructure:"max_completion_tokens"`
	Temperature         float32 `mapstructure:"temperature"`
}

// VectorDBConfig 包含向量数据库相关配置
type VectorDBConfig struct {
	Type           string `mapstructure:"type"`            // "inmemory" or "external"
	Dimension      int    `mapstructure:"dimension"`       // 维度
	SimilarityFunc string `mapstructure:"similarity_func"` // "cosine", "dot", "euclidean"
	MaxResults     int    `mapstructure:"max_results"`
}

// StorageConfig 持有存储配置
type StorageConfig struct {
	DocumentDir string `mapstructure:"document_dir"`
	DBPath      string `mapstructure:"db_path"`
}

// Load 从文件或环境变量中读取配置
func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("openai.embedding_model", "text-embedding-ada-002")
	viper.SetDefault("openai.completion_model", "gpt-3.5-turbo")
	viper.SetDefault("openai.max_completion_tokens", 1000)
	viper.SetDefault("openai.temperature", 0.7)
	viper.SetDefault("vectordb.type", "inmemory")
	viper.SetDefault("vectordb.dimension", 1536)
	viper.SetDefault("vectordb.similarity_func", "cosine")
	viper.SetDefault("vectordb.max_results", 5)
	viper.SetDefault("storage.document_dir", "./data/documents")
	viper.SetDefault("storage.db_path", "./data/gorag.db")

	if err := viper.ReadInConfig(); err != nil {
		// It's okay if config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	err := viper.Unmarshal(&config)
	return &config, err
}
