package main

import (
	"fmt"
    "log"
    "os"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/config"	
)

func main() {
	cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    st := &state{cfg: &cfg}

    cmds := &commands{handlers: make(map[string]func(*state, command) error)}
    cmds.register("login", handlerLogin)

    if len(os.Args) < 2 {
        fmt.Println("No command provided")
        os.Exit(1)
    }
    
    name := os.Args[1]
    args := os.Args[2:]

    cmd := command{name: name, arguments: args}

    if err := cmds.run(st, cmd); err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }
}
