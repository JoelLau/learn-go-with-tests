package concurrency

type Checker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(check Checker, websites []string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChannel := make(chan result)

	for _, website := range websites {
		go func(s string) {
			resultChannel <- result{s, check(s)}
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return
}
