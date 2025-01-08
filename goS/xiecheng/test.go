package main

import (
	"fmt"
	"time"
)

// 1. 配置结构体
type ServerConfig struct {
	Host     string
	Port     int
	Timeout  time.Duration
	Protocol string
}

// 2. Option 类型
type Option func(*ServerConfig)

// 3. Option 函数
func WithHost(host string) Option {
	return func(config *ServerConfig) {
		config.Host = host
	}
}

func WithPort(port int) Option {
	return func(config *ServerConfig) {
		config.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(config *ServerConfig) {
		config.Timeout = timeout
	}
}
func WithProtocol(protocol string) Option {
	return func(config *ServerConfig) {
		config.Protocol = protocol
	}
}

// 4. 配置函数
func NewServer(options ...Option) *ServerConfig {
	config := &ServerConfig{
		Host:     "localhost",
		Port:     8080,
		Timeout:  30 * time.Second,
		Protocol: "http",
	}

	for what, option := range options {
		option(config)
		fmt.Println(what, "==========", option)
	}
	return config
}

func main() {
	// 使用 Option 配置服务器
	server1 := NewServer()
	fmt.Printf("Server1: %+v\n", server1)

	server2 := NewServer(WithHost("127.0.0.1"), WithPort(9090))
	fmt.Printf("Server2: %+v\n", server2)

	server3 := NewServer(WithHost("example.com"), WithPort(8000), WithTimeout(1*time.Minute), WithProtocol("https"))
	fmt.Printf("Server3: %+v\n", server3)
}
