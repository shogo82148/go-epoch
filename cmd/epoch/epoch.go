package main

import (
	"fmt"
	"time"
)

func main() {
	var epoch int64
	fmt.Scanf("%d", &epoch)
	fmt.Println(time.Unix(epoch, 0))
}
