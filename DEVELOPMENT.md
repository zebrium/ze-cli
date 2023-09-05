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
To cross compile a binary locally, you can run make build.  This will leverage goreleaser to crosscompile the binary in all supported architectures

## Testing 
Ze has a suite of unit testing for internal modules to ensure business logic is accurate. You can run the full battery of tests with `make gotest` at the root of the module.

## Committing Changes
Before committing changes to ze-cli and opening a Pull Request, please run `make all` to ensure all checks are met in order
to make your PR pass the required checks. 

## Cleanup
In order to clean your environment after running make commands, you can run `make cleanup`.  
This will purge any internal directories created from previous make commands. 