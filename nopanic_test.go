package nopanic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	nopanic "github.com/gomatic/yze-go-nopanic"
)

func TestPanicCallIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nopanic.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, nopanic.Registration.Validate())
	assert.Equal(t, "yze/nopanic", nopanic.Registration.RuleID())
	assert.Same(t, nopanic.Analyzer, nopanic.Registration.Analyzer)
}
