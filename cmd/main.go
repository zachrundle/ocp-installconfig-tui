package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
	"github.com/zachrundle/ocp-installconfig-tui/internal"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("zts version %s\n", version.Get())
	}	
}