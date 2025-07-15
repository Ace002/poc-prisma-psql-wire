package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	wire "github.com/jeroenrinzema/psql-wire"

	"github.com/Ace002/poc-prisma-psql-wire/psql_wire_app/config"
	handlers "github.com/Ace002/poc-prisma-psql-wire/psql_wire_app/handlers"
)

func ListenAndServe(addr string, handler wire.ParseFn, cfg *config.Config) error {
	server, err := wire.NewServer(handler)
	if err != nil {
		return err
	}

	server.Auth = wire.ClearTextPassword(func(ctx context.Context, db, user, pass string) (context.Context, bool, error) {
		// log := logger.Get()

		if user != cfg.ExpectedUser {
			err := fmt.Errorf("unauthorized user: %s, but should be %s", user, cfg.ExpectedUser)
			log.Println("Auth failed", "error", err)
			return ctx, false, err
		}
		if pass != cfg.ExpectedPassword {
			err := fmt.Errorf("invalid password")
			log.Println("Auth failed", "error", err)
			return ctx, false, err
		}
		return ctx, true, nil
	})

	return server.ListenAndServe(addr)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "\n🔥 PANIC recovered: %v\n", r)
			fmt.Fprintf(os.Stderr, "🔥 STACK:\n%s\n", debug.Stack())

			// // Still try to log cleanly
			// if log := logger.Get(); log != nil {
			// 	log.Error("Fatal panic recovered", "error", r, "stack", string(debug.Stack()))
			// 	logger.Close()
			// }

			os.Exit(2)
		}
	}()

	// logPath := "/var/log/SPSE_Internal_Db/spse_internal_db.log"
	// if err := logger.Init(logPath); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Logger init failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer logger.Close()

	// log := logger.Get()
	log.Println("SPSE Internal DB proxy starting up...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	addr := fmt.Sprintf("0.0.0.0:%s", cfg.ListeningPort)
	log.Println("Starting server", "address", addr)

	err = ListenAndServe(addr, handlers.QueryHandler(cfg), cfg)
	if err != nil {
		log.Println("Server failed to start", "error", err, "stack", string(debug.Stack()))
		os.Exit(1)
	}
}
