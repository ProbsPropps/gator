package main

import (
	"fmt"

	"github.com/ProbsPropps/gator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	cfg.SetUser("Austin")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("DB URL: %s\nCurrent Username: %s\n", cfg.DBURL, cfg.CurrentUserName)
}
