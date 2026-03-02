package collector

import "time"

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		//TODO: run polling and data collection here
	}
}
