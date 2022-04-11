package conf

import (
  "log"
  "os"
  "strconv"
)

const (
  hostKey       = "RGB_HOST"
  portKey       = "RGB_PORT"
  jwtSecretKey  = "JWT_SECRET"
)

type Config struct {
  Host       string
  Port       string
  JwtSecret  string
}

func NewConfig() Config {
  host, ok := os.LookupEnv(hostKey)
  if !ok || host == "" {
    logAndPanic(hostKey)
  }

  port, ok := os.LookupEnv(portKey)
  if !ok || port == "" {
    if _, err := strconv.Atoi(port); err != nil {
      logAndPanic(portKey)
    }
  }
  jwtSecret, ok := os.LookupEnv(jwtSecretKey)
  if !ok || jwtSecret == "" {
    logAndPanic(jwtSecretKey)
  }

  return Config{
    Host:       host,
    Port:       port,
	  JwtSecret:  jwtSecret,
  }
}

func logAndPanic(envVar string) {
  log.Println("ENV variable not set or value not valid: ", envVar)
  panic(envVar)
}