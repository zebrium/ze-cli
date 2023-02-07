# ze cli Development Guide
The ze go cli was developed leveraging the [Cobra](https://github.com/spf13/cobra) go package. 

## Adding a new submodule
We leverage the cobra cli tool for adding new submodules and set up the default
templating for each command.  Usage instructions can be found [here](https://github.com/spf13/cobra-cli/blob/main/README.md) 

To add the incident subcommand  
`cobra-cli add incident `

Once that completed, we now want to create a read action, so we run:
`cobra-cli add list -p incidentCmd`

Since Cobra-cli adds the incident on a flat level, we want to rename the 
submodule commands to reflect the following structure 
`<submodule>_<action>.go`

## Common Folder
The common folder stores all common code for input sanitation and validation.
Functions in here can return a os.exit and interact with the std out and std err.
All other packages should return an object and an error and all system exits 
and std prompts should be handled in the cmd package.

## Viper Integration
This project leverages [Viper](https://github.com/spf13/viper) in order to support 
configuration injection either through a configuration file or through ENV variables. 
Viper configuration can be found [here](cmd/root.go)


## Building artifacts 
This project leverages github actions to build adn release the project based on github releases.  
To enable this, we use [goreleaser](https://github.com/goreleaser/goreleaser)

### Building Locally
Currently, this project does not use any special features of Go build. Artifacts for the current OS can be created using `go build` and if other OS or architectures are needed, you must export `GOOS` and `GOARCH` to the appropriate OS and CPU architecture.

## Testing 
Ze has a suite of unit testing for internal modules to ensure business logic is accurate. You can run the full battery of tests with `go test ./...` at the root of the module.