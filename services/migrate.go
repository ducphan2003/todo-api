package services

import (
	"flag"
	"fmt"
	// "todo-api/common"
	"todo-api/database"
	taskModel "todo-api/modules/task/taskmodel"
	userModel "todo-api/modules/user/usermodel"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

type MigrateCommand struct {
	fs        *flag.FlagSet
	tableName string
}

func NewMigrateCommand() *MigrateCommand {
	gc := &MigrateCommand{
		fs: flag.NewFlagSet("migrate", flag.ContinueOnError),
	}
	gc.fs.StringVar(&gc.tableName, "table", "all", "table name need to be migrated")
	return gc
}

func (g *MigrateCommand) Name() string {
	return g.fs.Name()
}

func (g *MigrateCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *MigrateCommand) Run() error {
	db := database.DB
	Logger.Infof("Migrate %s table...\n", g.tableName)
	if g.tableName == "all" {
		print("migrate all model")
		fmt.Println("-------------------", db)
		db.AutoMigrate(&userModel.User{})
		db.AutoMigrate(&taskModel.Task{})
	} else {
		print(g.tableName)
		switch g.tableName {
		case "users":
			db.AutoMigrate(&userModel.User{})
		case "tasks":
			db.AutoMigrate(&taskModel.Task{})
		}

	}
	Logger.Info("Migrate successfully")
	return nil
}
