package main

import (
	"github.com/tnkTaka/go-MorphologicalAnalysis/analysis"
	"fmt"
)

func main() {
	text := "golang で形態素解析を並列実行させるgolang"
	result := analysis.MorphologicalAnalysis(text)
	fmt.Println(result)
}
