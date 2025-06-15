package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/ProbsPropps/gator/internal/config"
	"github.com/ProbsPropps/gator/internal/database"
	_ "github.com/lib/pq"
)

func main(){
	var s state
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error - main: %v\n", err)
		os.Exit(1)
	}
	s.cfg = &cfg
	
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Error - main: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	s.db = dbQueries

	var cmds commands
	cmds.commandNames = make(map[string]func(*state, command) error)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	args := os.Args[1:]
	
	if len(args) < 1 {
		fmt.Println("Error - main: Not enough arguments")
		os.Exit(1)
	}

	if len(args) < 2 && !strings.Contains(args[0], "reset") {
		fmt.Println("Error - main: Need a username")
		os.Exit(1)
	}
	cmd := command{name: args[0], args: args}
	if err = cmds.run(&s, cmd); err != nil {
		fmt.Printf("Error - main: %v\n", err)
		os.Exit(1)
	}
}
