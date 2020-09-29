package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mneverov/comment/pkg/comment"
)

func main() {
	singlechecker.Main(comment.Analyzer)
}
