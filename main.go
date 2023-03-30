package main

import "os"

func main() {
	// bc := CreateBlockchain("1GPAnTbyFMJNL1onm7T7DUFGzJRChbFHfW", "3000")
	nodeID := os.Getenv("NODE_ID")
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
