// this is the module descriptor of this module created with `go mod init goilona`
// dependency management other project properties will be managed in here
// when working with dependencies, the first run of this project will create a "go.sum"
// containing the checksums of the dependnecis in use for this project
// both go.mod and go.sum should be pushed to the source control system

// the `go mod tidy` command will tidy up this verifying and updating the go.sum file.

// the `go mod vendor` command will create a "vendor" folder in the project root,
// containg all the third party dependencies including the transitive depdencies,
// this will allow enabling vendoring which uses the vendor folder as a source for the dependencis isntead of the internet.

// a dependnecy in go.mod file might look like this:
// require github.com/operator-framework/operator-lib v0.10.0
// and in go.sum you'll have the checksum of this specific version

// ref: https://go.dev/doc/modules/gomod-ref
module lets-go-go

go 1.17
