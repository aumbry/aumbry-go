package aumbry

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/ghodss/yaml"
)

// Aumbry Package Constants
const (
	YamlFile = "YamlFile"
)

// Aumbry is the core loader type.
type Aumbry struct {
	cfgType string
	model   interface{}
	rawData []byte
	options map[string]string
}

// New creates a Aumbry loader
func New(cfgType string, model interface{}, options map[string]string) *Aumbry {
	inst := new(Aumbry)
	inst.cfgType = cfgType
	inst.model = model
	inst.options = options

	return inst
}

// Load fetches the desired config and populates the model
func (a *Aumbry) Load() interface{} {
	switch a.cfgType {
	case YamlFile:
		return loadYaml(a)

	default:
		panic("Invalid config type")

	}
}

func loadYaml(a *Aumbry) interface{} {
	filename := a.options["CONFIG_FILENAME"]
	searchPaths := strings.Split(a.options["CONFIG_SEARCH_PATH"], ";")

	a.rawData = loadFile(filename, searchPaths)
	yaml.Unmarshal(a.rawData, &a.model)

	return a.model
}

func loadFile(name string, searchPaths []string) []byte {
	for _, searchPath := range searchPaths {
		fullPath := path.Join(searchPath, name)
		raw, err := ioutil.ReadFile(fullPath)

		if err == nil {
			return raw
		}
	}

	panic(fmt.Errorf("Couldn't find %s in any of the search paths", name))
}
