package main

import "github.com/go-MorphologicalAnalysis/analysis"

func main() {
	text := "golang で形態素解析を並列実行させる"
	analysis.MorphologicalAnalysis(text)
}
