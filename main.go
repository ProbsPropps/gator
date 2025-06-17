package main

import (
	"database/sql"
	"fmt"
	"os"
	"slices"

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
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	Args := os.Args[1:]
	
	baseCommands := []string{"reset", "users", "agg", "feeds", "following", "browse"}

	if len(Args) < 1 {
		fmt.Println("Error - main: Not enough arguments")
		os.Exit(1)
	}

	if len(Args) < 2 && !slices.Contains(baseCommands, Args[0]) {
		fmt.Println("Error - main: Need a username")
		os.Exit(1)
	}

	cmd := command{Name: Args[0], Args: Args}
	if err = cmds.run(&s, cmd); err != nil {
		fmt.Printf("Error - main: %v\n", err)
		os.Exit(1)
	}
}
