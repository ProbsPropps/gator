package main

import "fmt"

type command struct {
	name 	string
	args	[]string
}

type commands struct {
	commandNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f := c.commandNames[cmd.name]
	err := f(s, cmd)
	if err != nil {
		return fmt.Errorf("Error - run: %v", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandNames[name] = f
}
