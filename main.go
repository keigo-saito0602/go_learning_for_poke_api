package main

import (
	"os"

	"github.com/keigo-saito0602/go_learning_for_poke_api/cmd"
)

// @title go learning for poke api API
// @version 1.0
// @description POKE APIのドキュメント
// @BasePath /
func main() {
	os.Exit(cmd.Run())
}
