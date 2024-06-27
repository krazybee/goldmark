package goldmark_test

import (
	"testing"

	. "github.com/krazybee/goldmark"
	"github.com/krazybee/goldmark/parser"
	"github.com/krazybee/goldmark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
