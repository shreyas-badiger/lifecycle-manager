package version

import (
	"fmt"
	"runtime"
)

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// The main version number that is being run at the moment.
const Version = "0.6.1"

var BuildDate = ""

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
