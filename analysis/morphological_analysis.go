package analysis

import (
	"bufio"
	"io"
	"sync"

	"github.com/ikawaha/kagome/splitter"
	"github.com/ikawaha/kagome/tokenizer"
	"strings"

	"fmt"
)

func MorphologicalAnalysis(text string) map[string]int{
	ch := make(chan string, 1024)

	r := strings.NewReader(text)
	go nounFilter(ch, r)

	m := map[string]int{}
	for {
		s, ok := <-ch
		if !ok {
			break
		}
		m[s]++
	}

	fmt.Println(m)
	return m
}

func nounFilter(ch chan<- string, r io.Reader) {
	var wg sync.WaitGroup
	t := tokenizer.New()
	scanner := bufio.NewScanner(r)
	scanner.Split(splitter.ScanSentences)
	for scanner.Scan() {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			tokens := t.Tokenize(s)
			for _, tok := range tokens {
				if tok.Pos() == "名詞" {
					ch <- tok.Surface
				}
			}
		}(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		close(ch)
	}
	wg.Wait()
	close(ch)
}