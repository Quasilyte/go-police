package police

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var (
	configOnce sync.Once
	config     struct {
		AltImport map[string]string
	}
)

func requireConfig() {
	configOnce.Do(readConfig)
}

func readConfig() {
	source := "$GOPOLICE_CONFIG"
	filename := os.Getenv("GOPOLICE_CONFIG")
	if filename == "" {
		source = "default value, use $GOPOLICE_CONFIG to override it"
		filename = ".gopolice.json"
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("reading config: %v (file selected by the %s)",
			err, source))
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(fmt.Sprintf("decoding config: %v", err))
	}
}
