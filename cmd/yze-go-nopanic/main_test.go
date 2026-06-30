package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis"

	nopanic "github.com/gomatic/yze-go-nopanic"
)

func TestMainRunsTheAnalyzer(t *testing.T) {
	original := run
	t.Cleanup(func() { run = original })

	var got *analysis.Analyzer
	run = func(a *analysis.Analyzer) { got = a }

	main()

	require.NotNil(t, got)
	assert.Same(t, nopanic.Analyzer, got)
}
