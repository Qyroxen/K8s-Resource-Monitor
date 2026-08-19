package main

import (
	"fmt"
	"os"
)

// k8s_resource_monitor - Monitor K8s resource usage
func k8s_resource_monitor(path string) {
	fmt.Println("========================================")
	fmt.Println("  K8s-Resource-Monitor")
	fmt.Println("  Monitor K8s resource usage")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	k8s_resource_monitor(path)
}
