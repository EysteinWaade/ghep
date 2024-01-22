package slack

import (
	"os"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
	"github.com/navikt/ghep/internal/github"
)

func TestCreateCommitMessage(t *testing.T) {
	var err error
	commitTmpl, err = template.ParseFS(templates, "templates/commit.tmpl")
	if err != nil {
		t.Error(err)
	}

	want, err := os.ReadFile("goldenfiles/commit.json")
	if err != nil {
		t.Error(err)
	}

	event := github.CommitEvent{
		Compare: "https://github.com/test/compare/2d7f6c9...d6f21c8",
		Repository: github.Repository{
			Name:    "test",
			HtmlUrl: "https://github.com/test",
		},
		Commits: []github.Commit{
			{
				Id:      "d6f21c84",
				Message: "test",
				Url:     "https://github.com/test",
				Author: github.Author{
					Name:  "Ola Nordmann",
					Email: "ole@nordmann.no",
				},
			},
		},
	}

	t.Run("test", func(t *testing.T) {
		got, err := CreateCommitMessage("#test", event)
		if err != nil {
			t.Error(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("payload mismatch (-want +got):\n%s", diff)
		}
	})
}
