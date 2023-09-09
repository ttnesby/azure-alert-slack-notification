package main

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/divider"
)

func main() {
	l := divider.New()

	fmt.Printf("%v", *l)
}
