package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/rickcollette/kvdbcli/cmd"
)

func replMode() {
	rl, err := readline.New("> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	fmt.Println("Entering interactive mode. Type 'quit' or 'exit' to leave.")

	for {
		line, err := rl.Readline()
		if err != nil { // Handle errors or EOF
			break
		}

		command := strings.TrimSpace(line)

		// Handle exit command
		if command == "exit" || command == "quit" {
			break
		}

		// Handle REPL commands
		args := strings.Split(command, " ")
		switch args[0] {
		case "version":
			cmd.Execute() // You could also create a handler if needed
		case "insert":
			if len(args) < 3 {
				fmt.Println("Usage: insert <key> <value>")
			} else {
				insertInteractive(args[1], args[2])
			}
		case "read":
			if len(args) < 2 {
				fmt.Println("Usage: read <key>")
			} else {
				readInteractive(args[1])
			}
		case "update":
			if len(args) < 3 {
				fmt.Println("Usage: update <key> <new_value>")
			} else {
				updateInteractive(args[1], args[2])
			}
		case "delete":
			if len(args) < 2 {
				fmt.Println("Usage: delete <key>")
			} else {
				deleteInteractive(args[1])
			}
		case "snapshot":
			snapshotInteractive()
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

func insertInteractive(key, value string) {
	btree, err := cmd.LoadBtree()
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	if err := cmd.InsertKey(btree, key, value); err != nil {
		log.Fatalf("Error inserting key: %v", err)
	}
	fmt.Printf("Inserted key: %s\n", key)
}

func readInteractive(key string) {
	btree, err := cmd.LoadBtree()
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	val, err := cmd.ReadKey(btree, key)
	if err != nil {
		log.Fatalf("Error reading key: %v", err)
	}
	fmt.Printf("Value for key '%s': %s\n", key, val)
}

func updateInteractive(key, value string) {
	btree, err := cmd.LoadBtree()
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	if err := cmd.UpdateKey(btree, key, value); err != nil {
		log.Fatalf("Error updating key: %v", err)
	}
	fmt.Printf("Updated key: %s\n", key)
}

func deleteInteractive(key string) {
	btree, err := cmd.LoadBtree()
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	if err := cmd.DeleteKey(btree, key); err != nil {
		log.Fatalf("Error deleting key: %v", err)
	}
	fmt.Printf("Deleted key: %s\n", key)
}

func snapshotInteractive() {
	btree, err := cmd.LoadBtree()
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	if err := cmd.SnapshotBtree(btree); err != nil {
		log.Fatalf("Error taking snapshot: %v", err)
	}
	fmt.Println("Snapshot taken.")
}

func main() {
	if len(os.Args) == 1 {
		// No arguments passed, enter interactive mode
		replMode()
	} else {
		// Arguments are passed, use the existing CLI method
		cmd.Execute()
	}
}
