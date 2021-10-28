package fetcher

import (
	"fmt"
	"strings"
	"sync"

	"github.com/umutphp/awesome-cli/internal/package/node"
	"github.com/umutphp/awesome-cli/internal/package/parser"
)

const DOWNLOAD_CONCURRENCY = 4

type Progress struct {
	Found uint64
	Crawled uint64
	Errors uint64
}

func FetchAllRepos(update chan Progress) (Progress, error) {
	fetched, err := FetchAwesomeRootRepo()
	if err != nil {
		return Progress{1, 0, 1}, err
	}

	root := parser.ParseIndex(fetched)
	urls := getAllChildrenURLs(root.GetChildren())

	progress := Progress{
		Found: uint64(1 + len(urls)),
		Crawled: 1,
		Errors: 0,
	}
	update <- progress

	var wg sync.WaitGroup
	queue := make(chan string)
	errors := make(chan error)

	for i := 0; i < 4; i++ {
		go fetchWorker(queue, errors, &wg)
	}

	allErrors := MultiError{[]error{}}
	go func () {
		for err := range errors {
			if err == nil {
				progress.Crawled = progress.Crawled + 1
			} else {
				progress.Errors = progress.Errors + 1
				allErrors.Errors = append(allErrors.Errors, err)
			}
			update <- progress
		}
		close(update)
	}()

	for _, url := range urls {
		queue <- url
	}

	close(queue)
	wg.Wait()
	close(errors)

	if len(allErrors.Errors) == 0 {
		return progress, nil
	}
	return progress, allErrors
}

func getAllChildrenURLs(children []node.Node) []string {
	urls := []string{}
	for _, child := range children {
		if child.GetURL() != "" {
			urls = append(urls, child.GetURL())
		}
		urls = append(urls, getAllChildrenURLs(child.GetChildren())...)
	}
	return urls
}

type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	errStrings := []string{}
	for _, err := range e.Errors {
		errStrings = append(errStrings, err.Error())
	}
	return strings.Join(errStrings, "\n")
}

type FetchError struct {
	URL string
	Wrapped error
}

func (e FetchError) Unwrap() error {
	return e.Wrapped
}

func (e FetchError) Error() string {
	return fmt.Sprintf("failed to fetch %q: %v", e.URL, e.Wrapped)
}

func fetchWorker(queue chan string, errors chan error, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for url := range queue {
		if _, err := FetchAwesomeRepo(url); err != nil {
			errors <- FetchError{url, err}
		} else {
			errors <- nil
		}
	}
}