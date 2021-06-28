package reporting

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jtmelton/taw/domain"
)

func check(e error) {
	if e != nil {
		log.Fatalf(e.Error())
	}
}

// Reporter is a marker interface
type Reporter interface {
	write(extensions domain.ExtensionCounts, _options domain.Options)
}

// JSONReporter is an interface impl for json reporting
type JSONReporter struct{}

func (r JSONReporter) write(extensions domain.ExtensionCounts, _options domain.Options) {

	extensionCountsJSON, jsonErr := json.Marshal(extensions)
	check(jsonErr)

	var results []byte

	// populate results with either empty json array or the actual counts
	if extensions.ExtensionCounts == nil || len(extensions.ExtensionCounts) == 0 {
		results = []byte("{ \"extension-counts\": []}")
	} else {
		results = extensionCountsJSON
	}

	err := ioutil.WriteFile(_options.OutputFile, results, 0644)
	check(err)
}

// WriteReport is the generic function called by an outside class to serialize the results.
// Could support more formats in the future
func WriteReport(extensions domain.ExtensionCounts, _options domain.Options) {
	reporter := JSONReporter{}

	reporter.write(extensions, _options)
}
