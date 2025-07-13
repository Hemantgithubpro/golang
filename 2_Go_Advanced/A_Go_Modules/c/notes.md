# Steps
- Create a new module: `go mod init example.com/hello` or `go mod init github.com/user/repo`
- Add a dependency: `go get github.com/some/dependency`
- Update dependencies: `go get -u`
- Tidy up the module: `go mod tidy`
- View module information: `go list -m all`
- Remove unused dependencies: `go mod tidy`
- Check for vulnerabilities: `go mod verify`
- View the current module's dependencies: `go mod graph`
- Test a module: `go test ./...`