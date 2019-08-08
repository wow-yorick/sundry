package search

import "log"

type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search: Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

// Bing search will be added to this search session if this option
// is provided.
func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

// Yahoo search will be enabled if this option is provided as an argument
// to Submit.
func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

func OnlyFirst(s *searchSession) { s.first = true }

func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result

	for search := 0; search < len(session.searchers); search++ {
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		log.Println("search : submit : info : waiting for results...")
		result := <-session.resultChan

		log.Printf("search : submit : info : results used : results[%d]\n", len(result))
		results = append(results, result...)
	}

	log.Printf("search : submit : completed : found [%d] results\n", len(results))
	return results
}
