package processing

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jtmelton/taw/domain"
)

// Walk walks the directory tree, forking a new worker goroutine for the number of workers
func Walk(rootPath string, _options domain.Options) domain.ExtensionCounts {

	var extensionsMap = make(map[string]int)

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failure walking %s: %s", rootPath, err)
		}

		if info.IsDir() {
			return nil
		}

		extension := filepath.Ext(path)

		currentCount, exists := extensionsMap[extension]
		if exists {
			extensionsMap[extension] = currentCount + 1
		} else {
			extensionsMap[extension] = 1
		}

		return nil
	})

	var extensions []domain.ExtensionCount
	for ext, cnt := range extensionsMap {
		extensions = append(extensions,
			domain.ExtensionCount{Extension: ext, Count: cnt})
	}

	return domain.ExtensionCounts{ExtensionCounts: extensions}
}
