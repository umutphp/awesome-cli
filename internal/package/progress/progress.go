package progress

import (
	"fmt"

	"github.com/umutphp/awesome-cli/internal/package/fetcher"
)

func ProgressBar(done chan struct{}, updates chan fetcher.Progress) {
	for progress := range updates {
		percent := 100.0 * float64(progress.Crawled + progress.Errors) / float64(progress.Found)
		fmt.Print("\n\033[1A\033[K")
		fmt.Print("[")
		for i := 0; i < 20; i++ {
			if float64(i*5) <= percent {
				fmt.Print("=")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("] % 3.0f%%", percent)
	}
	close(done)
}