package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"
	"sync"
)

var (
	debug = flag.Bool("d", false, "Enable debug logging.")
	gi    = flag.Bool("i", false, "Search Google Images.")
	gv    = flag.Bool("v", false, "Search Google Videos.")
	gm    = flag.Bool("m", false, "Search Google Maps.")
	gw    = flag.Bool("w", true, "Search Google.")
	yt    = flag.Bool("yt", false, "Search YouTube.")
)

var urls = map[*bool]string{
	gi: "https://www.google.com/search?tbm=isch",
	gv: "https://www.google.com/search?tbm=vid",
	gw: "https://www.google.com/search",
	gm: "https://www.google.com/maps",
	yt: "https://www.youtube.com/results",
}

var qskeys = map[*bool]string{
	yt: "search_query",
	gi: "q",
	gv: "q",
	gm: "q",
	gw: "q",
}

var maxRoutines = len(urls)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Expected search query; got nothing.")
	}
	if *debug {
		log.Printf("Searching for: %v", args)
	}

	var (
		errc = make(chan error, maxRoutines)
		wg   sync.WaitGroup
	)

	query := strings.Join(args, " ")

	for _, flag := range []*bool{yt, gi, gv, gm, gw} {
		if !*flag {
			continue
		}
		wg.Add(1)
		go search(flag, query, errc, &wg)
	}

	wg.Wait()
	close(errc)

	for err := range errc {
		log.Printf("Error opening URL: %v", err)
	}
}

func search(flag *bool, query string, errc chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	var err error

	rawurl := urls[flag]
	endurl, err := getFinalURL(rawurl, qskeys[flag], query)

	if err != nil {
		errc <- fmt.Errorf("parsing url %s: %v", rawurl, err)
		return
	}
	if *debug {
		log.Printf("Using URL %v\n", *endurl)
	}

	//FIXME not portable and works for Darwin only
	cmd := exec.Command("open", *endurl)
	if err = cmd.Start(); err != nil {
		errc <- err
	}
}

func getFinalURL(rawurl, qskey string, query string) (*string, error) {
	p, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	qs := p.Query()
	qs.Set(qskey, query)
	qs.Encode()

	endurl := "https://" + p.Host + p.Path + "?" + qs.Encode()
	return &endurl, nil
}
