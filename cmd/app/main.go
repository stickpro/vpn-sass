package main

import (
	"fmt"
	"github.com/stickpro/vpn-sass/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		return
	}

	fmt.Println(cfg.HTTP)
}
