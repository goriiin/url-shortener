## содержит
- дефолтные локальные конфигурации сервера


### плохая конфигурация:

```go
package config

// Config - супер конфиг
type Config struct {
	Logger LoggerConfig
	Server ServerConfig
	Postgres PSConfig
}

type LoggerConfig struct {
	// код
}

type ServerConfig struct {
	// код
}

type PSConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int `json:"port" yaml:"port"`
	// код
}
```

### почему?
- появляется пакет, который тянет множество зависимостей
- риск циклических зависимостей

### хорошая конфигурация

пакет владеет своим конфигом

свои дефолт значения, ретрая, портов, таймаутов