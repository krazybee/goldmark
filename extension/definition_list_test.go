package extension

import (
	"testing"

	"github.com/krazybee/goldmark"
	"github.com/krazybee/goldmark/renderer/html"
	"github.com/krazybee/goldmark/testutil"
)

func TestDefinitionList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			DefinitionList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/definition_list.txt", t, testutil.ParseCliCaseArg()...)
}
