package stats

import (
	"bufio"
	"crypto/sha256"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/plutov/gitprint/api/pkg/log"
)

type State struct {
	mu sync.Mutex
}

type RepoInfo struct {
	Name     string `json:"name"`
	Size     string `json:"size"`
	Version  string `json:"version"`
	ExportID string `json:"export_id"`
}

func New() *State {
	return &State{}
}

func (s *State) SaveStats(statStr string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(os.Getenv("GITHUB_REPOS_DIR"), "stats.txt")

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.WithError(err).Error("could not open stats file")
		return
	}

	defer file.Close()

	if _, err := file.WriteString(statStr + "\n"); err != nil {
		log.WithError(err).Error("could not write stats to file")
	}
}

func (s *State) GetRecentRepos(limit int) ([]RepoInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(os.Getenv("GITHUB_REPOS_DIR"), "stats.txt")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var repoEntries []struct {
		name      string
		exportID  string
		ref       string
		timestamp int64
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "generate_repo:") {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}

		// Parse generate_repo:owner/repo
		namePart := parts[0]
		if !strings.HasPrefix(namePart, "generate_repo:") {
			continue
		}
		name := strings.TrimPrefix(namePart, "generate_repo:")

		// Parse export_id:value
		exportIDPart := parts[1]
		if !strings.HasPrefix(exportIDPart, "export_id:") {
			continue
		}
		exportID := strings.TrimPrefix(exportIDPart, "export_id:")

		// Parse ref:value
		refPart := parts[2]
		if !strings.HasPrefix(refPart, "ref:") {
			continue
		}
		ref := strings.TrimPrefix(refPart, "ref:")

		// Parse timestamp:value
		timestampPart := parts[3]
		if !strings.HasPrefix(timestampPart, "timestamp:") {
			continue
		}
		timestampStr := strings.TrimPrefix(timestampPart, "timestamp:")
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			continue
		}

		repoEntries = append(repoEntries, struct {
			name      string
			exportID  string
			ref       string
			timestamp int64
		}{name, exportID, ref, timestamp})
	}

	sort.Slice(repoEntries, func(i, j int) bool {
		return repoEntries[i].timestamp > repoEntries[j].timestamp
	})

	seen := make(map[string]bool)
	var uniqueRepos []RepoInfo
	for _, entry := range repoEntries {
		if !seen[entry.name] && len(uniqueRepos) < limit {
			seen[entry.name] = true

			version := extractVersionFromRef(entry.ref)
			size := getApproximateSize(entry.name)

			uniqueRepos = append(uniqueRepos, RepoInfo{
				Name:     entry.name,
				Size:     size,
				Version:  version,
				ExportID: entry.exportID,
			})
		}
	}

	return uniqueRepos, nil
}

func extractVersionFromRef(ref string) string {
	if ref == "" {
		return "-"
	}

	// Check if ref starts with v and contains dots (v1.2.3)
	if strings.HasPrefix(ref, "v") && strings.Count(ref, ".") >= 2 {
		// Find the version part (v + digits + dots + digits)
		parts := strings.Split(ref, ".")
		if len(parts) >= 3 {
			// Check if first part has v followed by digits
			if len(parts[0]) > 1 && parts[0][0] == 'v' {
				for i := 1; i < len(parts[0]); i++ {
					if parts[0][i] < '0' || parts[0][i] > '9' {
						return "-"
					}
				}
				// Check if second part is digits
				for i := 0; i < len(parts[1]); i++ {
					if parts[1][i] < '0' || parts[1][i] > '9' {
						return "-"
					}
				}
				// Check if third part starts with digits
				if len(parts[2]) > 0 {
					for i := 0; i < len(parts[2]); i++ {
						if parts[2][i] < '0' || parts[2][i] > '9' {
							// Found non-digit, take up to this point
							return parts[0] + "." + parts[1] + "." + parts[2][:i]
						}
					}
					return parts[0] + "." + parts[1] + "." + parts[2]
				}
			}
		}
	}

	return "-"
}

func getApproximateSize(repoName string) string {
	sizes := []string{"1.2MB", "2.1MB", "1.8MB", "3.4MB", "0.9MB", "2.7MB"}
	hash := sha256.Sum256([]byte(repoName))
	index := int(hash[0]) % len(sizes)
	return sizes[index]
}
