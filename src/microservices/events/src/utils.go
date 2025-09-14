package main

import (
	"fmt"
)

func getTopic(entity string) string {
	return fmt.Sprintf("%s-events", entity)
}
