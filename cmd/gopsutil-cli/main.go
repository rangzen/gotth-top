package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	count, _ := cpu.Counts(true)
	fmt.Println("CPU Count:", count)
	infoStats, _ := cpu.Info()
	for _, info := range infoStats {
		fmt.Println(info)
	}
}