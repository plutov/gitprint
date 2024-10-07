package builder

import (
	"testing"

	"github.com/google/go-github/v65/github"
)

func TestGenerateDocument(t *testing.T) {
	repo := &github.Repository{
		FullName:        github.String("testdata/testrepo"),
		Description:     github.String("test description"),
		StargazersCount: github.Int(10000),
		ForksCount:      github.Int(500),
		License: &github.License{
			Name: github.String("MIT"),
		},
	}
	contributors := []*github.Contributor{
		{
			Login:         github.String("plutov"),
			AvatarURL:     github.String("https://avatars.githubusercontent.com/u/1124859?v=4"),
			Contributions: github.Int(100),
		},
	}

	tests := []struct {
		repository    *github.Repository
		contributors  []*github.Contributor
		outputDir     string
		isNilErr      bool
		nodesCount    int
		chaptersCount int
		rootReadme    bool
	}{
		{repo, contributors, "notfound", false, 0, 0, false},
		{repo, contributors, "./testdata/testrepo", true, 19, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.outputDir, func(t *testing.T) {
			doc, err := GenerateDocument(tt.repository, tt.contributors, "a4d5gq", tt.outputDir)
			if tt.isNilErr && err != nil {
				t.Errorf("expecting nil error, got %v", err)
			}
			if doc != nil {
				if len(doc.Nodes) != tt.nodesCount {
					t.Errorf("expecting %d nodes, got %d", tt.nodesCount, len(doc.Nodes))
				}

				chaptersCount := 0
				for i, node := range doc.Nodes {
					if node.Type == NodeTypeChapter {
						// validate root chapter
						if chaptersCount == 0 {
							content := node.Content.(ContentChapter)
							if node.Title != "root" {
								t.Errorf("expecting root title, got %s", node.Title)
							}
							if content.VersionMajor != 0 || content.VersionMinor != 0 || content.VersionPatch != 0 {
								t.Errorf("expecting 0.0.0 version, got %d.%d.%d", content.VersionMajor, content.VersionMinor, content.VersionPatch)
							}

							// next node should be root readme
							if tt.rootReadme {
								nextNode := doc.Nodes[i+1]
								if nextNode.Title != "README.md" {
									t.Errorf("expecting README.md title, got %s", nextNode.Title)
								}
							}
						}
						chaptersCount++
					}
				}

				if chaptersCount != tt.chaptersCount {
					t.Errorf("expecting %d chapters, got %d", tt.chaptersCount, chaptersCount)
				}
			}
		})
	}
}
