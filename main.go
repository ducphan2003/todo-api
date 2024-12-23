package main

import (
	"log"
	"os"
	"todo-api/database"
	"todo-api/services"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	services.InitializeLogger()
	initialDB()
	services.Logger.Infof("%s", os.Args)
	if len(os.Args) > 1 {
		handleRunScript()
		return
	}
	services.StartGRPCServer(":50051")
}

func initialDB() {
	var err error
	database.DB, err = database.GormOpen()
	if err != nil {
		services.Logger.Panicf("fatal error in database file: %s", err)
	}
}

func handleRunScript() {
	cmds := []services.Runner{
		services.NewMigrateCommand(),
	}
	subcommand := os.Args[1]
	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			if err := cmd.Run(); err != nil {
				services.Logger.Error(err)
			}
			break
		}
	}
}
