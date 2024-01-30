# ZE CLI Tool

`ze` is Zebrium's command line interface for uploading log events from files or streams.  Please visit the official [ScienceLogic's docs page](https://docs.sciencelogic.com/zebrium/latest/Content/Web_Zebrium/03_Log_Collectors_Uploads/File_Uploads_ze.html) for the full documentation.

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

 The ze cli tool does support setting global variables in a .ze file for easy
 configuration. The default location of this is `$HOME/.ze`, however this can overridden
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

## Upload Commands

The ze up command is used to upload log event data to your Zebrium instance from a file or stream (STDIN) with appropriate metadata.  To see a full list of upload options, please run the command:

```bash
ze up --help
```

Example Usages:

Ingest three log files associated with the same support case "sr12345" (does not assume a .ze configuration file exists):

```ze up --file=/casefiles/sr12345/messages.log --svcgrp=sr12345 --host=node01 --log=messages --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>```

```ze up --file=/casefiles/sr12345/application.log --svcgrp=sr12345 --host=node01 --log=application --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>```

```ze up --file=/casefiles/sr12345/db.log --svcgrp=sr12345 --host=db01 --log=db --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>```

Ingest a continuous tail of /var/log/messages. When reading from a stream, such as STDIN, rather than from a file, ze requires the –log flag (assumes a .ze configuration file exists):

```tail -f /var/log/messages | ze up --log=varlogmsgs --svcgrp=monitor01 --host=mydbhost```

## Batch Commands

Please see [zebrium batch documentation](ze_batch_uploads.md)

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
