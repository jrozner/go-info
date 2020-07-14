package info

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	// version is the version of the software
	version string

	// commitHash is the git commit hash that has been built
	commitHash string

	// runtime is the version of Go that was used to build the software
	runtimeVersion string = runtime.Version()

	// buildDate is the time the software was built
	buildDate string

	cmdOpt bool
)

func init() {
	flag.BoolVar(&cmdOpt, "version", false, "output the software version and build information")
}

// Data is a struct to contain the build and version information
type Data struct {
	Version    string
	CommitHash string
	Runtime    string
	BuildDate  time.Time
}

// Values returns a Data struct containing the stored build and version information
func Values() (*Data, error) {
	date, err := time.Parse(time.RFC1123, buildDate)
	if err != nil {
		return nil, err
	}

	return &Data{
		Version:    version,
		CommitHash: commitHash,
		Runtime:    runtimeVersion,
		BuildDate:  date,
	}, nil
}

// Print prints the build and version information
func Print() {
	fmt.Printf("Version: %s\nCommit Hash: %s\nGo Runtime: %s\nBuild Date: %s\n", version, commitHash, runtimeVersion, buildDate)
}

// Flag returns true if the version flag has been set and false if it has not
func Flag() bool {
	return cmdOpt
}
