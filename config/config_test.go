package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// A helper function to replace log.Fatal in tests
func replaceLogFatal(f func(format string, v ...interface{})) func() {
	old := logFatal
	logFatal = f
	return func() { logFatal = old }
}

func TestLoadConfigMissing(t *testing.T) {
	// Capture the log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr) // Restore original log output

	// Replace logFatal with a custom function
	restoreLogFatal := replaceLogFatal(func(format string, v ...interface{}) {
		_, err := fmt.Fprintf(&buf, format, v...)
		assert.NoError(t, err)
	})

	// Ensure logFatal is restored after the test
	defer restoreLogFatal()

	// Call the main function
	_, _ = LoadConfig("aaa.yaml")

	// Check the log output
	expected := "Error reading config file: Config File \"aaa.yaml\" Not Found in"
	if !strings.HasPrefix(buf.String(), expected) {
		t.Errorf("expected %q, but got %q", expected, buf.String())
	}
}

func TestLoadConfig(t *testing.T) {
	viper.AutomaticEnv()

	config, err := LoadConfig("config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, int32(100), config.Frequency)
	// HttpServer
	assert.Equal(t, 8088, config.HttpServer.Port)
	assert.Equal(t, "debug", config.HttpServer.Mode)
	assert.Equal(t, "/", config.HttpServer.ContextPath)
	// Storage
	assert.Equal(t, "localhost:6379", config.Storage.RedisHost)
	assert.Equal(t, 0, config.Storage.RedisDb)
	assert.Equal(t, "Scheduler", config.Storage.SchedulerChanel)
	assert.Equal(t, "TaskExecution", config.Storage.TaskExecutionChanel)
	assert.Equal(t, "scheduled_tasks", config.Storage.SchedulerKeyName)
	// Kafka
	//assert.Equal(t, true, config.Kafka.Enabled)
	assert.Equal(t, "localhost:9092", config.Kafka.Brokers)
	assert.Equal(t, "scheduler", config.Kafka.Group)
	assert.Equal(t, "Scheduler", config.Kafka.SchedulerTopic)
	assert.Equal(t, "TaskExecution", config.Kafka.TaskExecutionTopic)
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Set environment variable
	err := os.Setenv("REDIS_HOST", "XXX")
	assert.NoError(t, err)
	err = os.Setenv("REDIS_PASS", "YYY")
	assert.NoError(t, err)
	err = os.Setenv("BROKERS", "ZZZ")
	assert.NoError(t, err)

	// Automatically read environment variables
	viper.AutomaticEnv()

	// Bind the environment variable to the config key
	err = viper.BindEnv("storage.redisHost", "REDIS_HOST")
	assert.NoError(t, err)
	err = viper.BindEnv("storage.redispass", "REDIS_PASS")
	assert.NoError(t, err)
	err = viper.BindEnv("kafka.brokers", "BROKERS")
	assert.NoError(t, err)

	config, err := LoadConfig("config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, 8088, config.HttpServer.Port)
	// should be overridden
	assert.Equal(t, "XXX", config.Storage.RedisHost)
	assert.Equal(t, "YYY", config.Storage.RedisPass)
	assert.Equal(t, "ZZZ", config.Kafka.Brokers)
}
