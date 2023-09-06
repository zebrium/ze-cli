# ZE CLI Tool

`ze` is Zebrium's command line interface for uploading log events from files or streams.

## Getting Started

### Installing

* Download the corresponding release from the releases [page](https://github.com/zebrium/ze-cli/releases)
* Unzip the downloaded file in
* Set up your path in your shell config to include the new binary
* Start a new terminal and test your installation 
 `ze -v`

## Configuration

The ze cli tool supports a variety of ways to set its parameters.  All parameters are 
settable via args.  To find out the args available and required for each call, use `ze -help` 
or `ze <subcommand> -help`  When leveraging the configuration file or ENV variables, ze cli will use the following 
precedence: Config File -> Env Files -> Command Line Args

### Configuration File

 The ze cli tool does support setting global variables in a .ze.yaml file for easy 
 configuration. The default location of this is `$HOME/.ze`, however this can overriden
 with passing a new path with the `--config` option. The contents of that file is as follows:

``` bash
auth: XXXXXXXXXX
url: https://cloud.zebrium.com
```

### Environment Variables

The ze cli supports setting the following env variables 

``` bash
ZE_AUTH: XXXXXXXXXXXX
ZE_URL: https://cloud.zebrium.com
```

## Usage

Use `ze -help` for a complete list of command options

## Migrating from the perl based ze-cli

The existing perl based application can be found [here](/legacy/bin)

### .zerc file

 The .zerc file is now replaced with a .ze file that accepts the configuration
 in yaml.  This is described [here](#configuration-file)  This means that configs that was specified as

```text
url=<ZE_LOG_COLLECTOR_URL>
auth=<ZE_LOG_COLLECTOR_TOKEN>
```

will now need to be 

```yaml
url: <ZE_LOG_COLLECTOR_URL>
auth: <ZE_LOG_COLLECTOR_TOKEN>
```


### ENV Variables

We now support setting env variables. currently we support the following list: 

```text
ZE_URL = <ZE_LOG_COLLECTOR_URL>
ZE_AUTH = <ZE_LOG_COLLECTOR_TOKEN>
```

## Contributors

* Braeden Earp (ScienceLogic)
