package sarif

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/owenrumney/go-sarif/models"
)

// Version is the version of Sarif to use
type Version string

// Version210 represents Version210 of Sarif
const Version210 Version = "2.1.0"

var versions = map[Version]string{
	Version210: "http://json.schemastore.org/sarif-2.1.0-rtm.4",
}

// Report is the encapsulating type representing a Sarif Report
type Report struct {
	Version string        `json:"version"`
	Schema  string        `json:"$schema"`
	Runs    []*models.Run `json:"runs"`
}

// New Creates a new Report or returns an error
func New(version Version) (*Report, error) {
	schema, err := getVersionSchema(version)
	if err != nil {
		return nil, err
	}
	return &Report{
		Version: string(version),
		Schema:  schema,
		Runs:    []*models.Run{},
	}, nil
}

// AddRun allows adding run information to the current report
func (sarif *Report) AddRun(toolName, informationURI string) *models.Run {
	run := models.NewRun(toolName, informationURI)
	sarif.Runs = append(sarif.Runs, run)
	return run
}

func getVersionSchema(version Version) (string, error) {
	for ver, schema := range versions {
		if ver == version {
			return schema, nil
		}
	}
	return "", fmt.Errorf("version [%s] is not supported", version)
}

// Write writes the JSON as a string with no formatting
func (sarif *Report) Write(w io.Writer) error {
	marshal, err := json.Marshal(sarif)
	if err != nil {
		return err
	}
	_, err = w.Write(marshal)
	return err
}

// PrettyWrite writes the JSON output with indentation
func (sarif *Report) PrettyWrite(w io.Writer) error {
	marshal, err := json.MarshalIndent(sarif, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(marshal)
	return err
}
