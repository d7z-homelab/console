package main

import (
  "context"
  "flag"
  "github.com/spf13/afero"
  "go.uber.org/zap"
  "os"
  "time"
  "xorm.io/xorm"
)

import (
  _ "modernc.org/sqlite"
)

var (
  bind     = ""
  database = ""
  storage  = ""
)

func init() {

  flag.StringVar(&bind, "bind", ":8080", "bind address")
  flag.StringVar(&database, "database", "", "database address")
  flag.StringVar(&storage, "storage", "", "storage dir")
}

func main() {
  flag.Parse()
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  zap.ReplaceGlobals(logger)
  engine, err := xorm.NewEngine("sqlite", database)
  if err != nil {
    zap.L().Fatal("failed to connect database", zap.Error(err))
  }
  _, err = engine.Exec("select 1")
  if err != nil {
    zap.L().Fatal("failed to connect database", zap.Error(err))
  }

  defer engine.Close()
  fs := afero.NewBasePathFs(afero.NewOsFs(), storage)
  if err = fs.MkdirAll("", 0755); err != nil && !os.IsExist(err) {
    zap.L().Fatal("failed to create directory", zap.Error(err))
  }
  ctx := context.Background()
  ctx = context.WithValue(ctx, "db", engine)
  ctx = context.WithValue(ctx, "fs", fs)
  //router := chi.NewRouter()
  time.Sleep(10000)
}
