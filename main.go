package main
/* 
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" down
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
*/
import (
	"fmt"
    "log"
    "os"
    "database/sql"

    _ "github.com/lib/pq"

	"github.com/Karina-Pogorzelec/blog_aggregator/internal/config"
    "github.com/Karina-Pogorzelec/blog_aggregator/internal/database"	
)


type state struct {
	db  *database.Queries
	cfg	*config.Config
}

func main() {
	cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    db, err := sql.Open("postgres", cfg.DBURL)
    if err != nil {
        log.Fatalf("error connecting to database: %v", err)
    }

    dbQueries := database.New(db)

    st := &state{cfg: &cfg, db: dbQueries}

    cmds := &commands{handlers: make(map[string]func(*state, command) error)}

    cmds.register("login", handlerLogin)
    cmds.register("register", handlerRegister)
    cmds.register("reset", handlerReset)
    cmds.register("users", handlerUsers)
    cmds.register("agg", handlerAgg)
    cmds.register("addfeed", handlerAddFeed)

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
