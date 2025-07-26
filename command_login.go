package main

import (
	"fmt"
	"github.com/mcbk51/blog_aggregator/internal/config"
)

type state struct{
 *config.Config	
}

type command struct{
	name string
  args []string
}

func handlerLogin(s *state, cmd command) error{
	if len(cmd.args) == 0 {
		return fmt.Errorf("The login handler expects a single argument, the username")
	} 
  
	username := cmd.args[0]
	s.user = username

	fmt.Print("The user has been set: %s\n", username)
  return nil	
}

type commands struct{
  handlers	map[string]func(*state, command) error 
}

func (c *commands) run(s *state, cmd command) error {
  handler, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("command '%s' not found", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error){
	if c.handlers == nil {
	  c.handlers = make(map[string]fuc(*state, command) error)	
	}
	c.handlers[name] = f
}

func handlerLogin(s *state, cmd command) error{
	if len(cmd.args) == O{
		return fmt.Errorf(" The login handler expects a single argument")
	}

	username := cmd.args[0]
	s.User = username
	fmt.Printf("The user has been set: %s\n", username)
	return nil
}
