package processing

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jtmelton/taw/domain"
)

// Walk walks the directory tree, forking a new worker goroutine for the number of workers
func Walk(rootPath string, _options domain.Options) domain.ExtensionCounts {

	var extensionCountsMap = make(map[string]int)
	var extensionBytesMap = make(map[string]int)

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failure walking %s: %s", rootPath, err)
		}

		if info.IsDir() {
			return nil
		}

		extension := filepath.Ext(path)
		stats, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}
		sizeInBytes := stats.Size()

		currentCount, exists := extensionCountsMap[extension]
		currentBytes := extensionBytesMap[extension]
		if exists {
			extensionCountsMap[extension] = currentCount + 1
			extensionBytesMap[extension] = currentBytes + int(sizeInBytes)
		} else {
			extensionCountsMap[extension] = 1
			extensionBytesMap[extension] = int(sizeInBytes)
		}

		return nil
	})

	var extensions []domain.ExtensionCount
	for ext, cnt := range extensionCountsMap {
		bytes := extensionBytesMap[ext]
		extensions = append(extensions,
			domain.ExtensionCount{Extension: ext, Count: cnt, Bytes: bytes})
	}

	return domain.ExtensionCounts{ExtensionCounts: extensions}
}
