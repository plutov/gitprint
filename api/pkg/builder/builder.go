package builder

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v65/github"
	"github.com/plutov/gitprint/api/pkg/log"
)

const (
	NodeTypeMeta         = "meta"
	NodeTypeContributors = "contributors"
	NodeTypeChapter      = "chapter"
	NodeTypeFile         = "file"
)

type ContentMeta struct {
	FullName        string `json:"fullName"`
	Description     string `json:"description"`
	ForksCount      int    `json:"forksCount"`
	StargazersCount int    `json:"stargazersCount"`
	License         string `json:"license"`
	CommitSHA       string `json:"commitSHA"`
}

type Contributor struct {
	Username      string `json:"username"`
	AvatarURL     string `json:"avatarURL"`
	Contributions int    `json:"contributions"`
}

type ContentContributors struct {
	Contributors []Contributor `json:"contributors"`
}

type ContentChapter struct {
	VersionMajor int `json:"versionMajor"`
	VersionMinor int `json:"versionMinor"`
	VersionPatch int `json:"versionPatch"`
}

type ContentFile struct {
	Content []byte `json:"content"`
}

type DocumentNode struct {
	Type    string      `json:"type"`
	Title   string      `json:"title"`
	Content interface{} `json:"content"`
}

type Document struct {
	Nodes []DocumentNode
}

func GenerateDocument(repo *github.Repository, contributors []*github.Contributor, commitSHA string, outputDir string) (*Document, error) {
	logCtx := log.With("repo", repo.GetFullName(), "outputDir", outputDir)
	logCtx.Info("generating document")

	doc := &Document{
		Nodes: []DocumentNode{},
	}
	doc.Nodes = append(doc.Nodes, DocumentNode{
		Type:  NodeTypeMeta,
		Title: repo.GetFullName(),
		Content: ContentMeta{
			FullName:        repo.GetFullName(),
			Description:     repo.GetDescription(),
			ForksCount:      repo.GetForksCount(),
			StargazersCount: repo.GetStargazersCount(),
			License:         repo.GetLicense().GetName(),
			CommitSHA:       commitSHA,
		},
	})

	contributorsList := []Contributor{}
	for _, c := range contributors {
		contributorsList = append(contributorsList, Contributor{
			Username:      c.GetLogin(),
			AvatarURL:     c.GetAvatarURL(),
			Contributions: c.GetContributions(),
		})
	}
	doc.Nodes = append(doc.Nodes, DocumentNode{
		Type:  NodeTypeContributors,
		Title: "Top Contributors",
		Content: ContentContributors{
			Contributors: contributorsList,
		},
	})

	var (
		versionMajor int = 0
		versionMinor int = 0
		versionPatch int = 0
	)

	lastChapterIndex := -1
	err := filepath.WalkDir(outputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		path = filepath.Clean(path)
		path = strings.TrimPrefix(path, filepath.Clean(outputDir))
		path = strings.TrimPrefix(path, string(os.PathSeparator))

		isRootDir := path == ""
		if d.IsDir() {
			title := path
			if isRootDir {
				title = "root"
			} else {
				subfoldersCount := strings.Count(path, string(os.PathSeparator))
				if subfoldersCount == 1 {
					versionMajor++
					versionMinor = 0
					versionPatch = 0
				} else if subfoldersCount == 2 {
					versionMinor++
					versionPatch = 0
				} else {
					versionPatch++
				}
			}

			doc.Nodes = append(doc.Nodes, DocumentNode{
				Type:  NodeTypeChapter,
				Title: title,
				Content: ContentChapter{
					VersionMajor: versionMajor,
					VersionMinor: versionMinor,
					VersionPatch: versionPatch,
				},
			})

			lastChapterIndex = len(doc.Nodes) - 1
		} else {
			// read file contents
			f, err := os.ReadFile(filepath.Join(outputDir, path))
			if err != nil {
				logCtx.With("path", path).WithError(err).Error("unable to read file")
				return err
			}

			node := DocumentNode{
				Type:  NodeTypeFile,
				Title: path,
				Content: ContentFile{
					Content: f,
				},
			}

			// make sure README.md is the first node in the chapter
			if strings.ToLower(path) == "readme.md" {
				// insert after last chapter node
				doc.Nodes = append(doc.Nodes[:lastChapterIndex+1], append([]DocumentNode{node}, doc.Nodes[lastChapterIndex+1:]...)...)
			} else {
				doc.Nodes = append(doc.Nodes, node)
			}
		}

		return nil
	})
	if err != nil {
		logCtx.WithError(err).Error("unable to walk directories")
		return nil, err
	}

	logCtx.With("nodes", len(doc.Nodes)).Info("document generated")

	return doc, nil
}
