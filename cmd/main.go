package main

import (
	"fmt"

	"github.com/codescalersinternships/psutil-go-DohaElsawy/psutil"
)


func main() {
	c := psutil.InitCPU()
	c.GetCpuInfo()

	fmt.Printf(c.String())
}