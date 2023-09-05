# ZE DETAILS
`ze` is Zebrium's command line interface for uploading log events from files or streams.
## Features

### up (upload)
Upload log event data to your Zebrium instance from a file or stream (stdin) with appropriate meta data.

### help
Display help on ze command usage.

### help_adv
Display advanced help on ze command usage.

## Getting Started

### Prerequisites

* perl
* perl JSON module
* curl
* Collector token from Zebrium (available from the Log Collector Setup page in the Zebrium UI)
* URL to your instance of Zebrium (available from the Log Collector Setup page in the Zebrium UI)

### Installing

* Download `/bin/ze` from the Zebrium GitHub repository here: [https://github.com/zebrium/ze-cli](https://github.com/zebrium/ze-cli/tree/master/legacy)
* Move `bin/ze` to the appropriate bin directory in your PATH
* Ensure ze command is executable: `chmod 755 <path_to_ze_command>`

**Note:** ze requires the JSON Perl module.
* To install on Linux (Ubuntu)
```
sudo apt-get install libjson-perl
```
* To install on mac OS
```
brew install cpanm
sudo cpanm install JSON
```

## Configuration
No configuration is required. All options can be specified as command line arguments. However, see the **Setup** section below for information on configuring your .zerc file.

### Setup
For convenience, the collector TOKEN and URL can be specified in your $HOME/.zerc file.

Your ZE_LOG_COLLECTOR_URL and ZE_LOG_COLLECTOR_TOKEN are available in the the Zebrium UI under the Log Collector Setup page.

Example .zerc file:
```
url=<ZE_LOG_COLLECTOR_URL>
auth=<ZE_LOG_COLLECTOR_TOKEN>
```

### Environment Variables
None

## Usage
Use `ze help` for a complete list of command options.

### Command Syntax and Options
```
  ze up                                                                              \
    [--url=<url>] [--auth=<token>]                                                   \
    [--file=<path>] [--log=<logtype>] [--host=<hostname>] [--svcgrp=<service-group>]

    --url      - Zebrium Log Collector URL <ZE_LOG_COLLECTOR_URL> (omit to look for url=<url> line in $HOME/.zerc)
    --auth     - Zebrium Log Collector Token <ZE_LOG_COLLECTOR_TOKEN> (omit to look for auth=<token> line in $HOME/.zerc)
    --file     - Path to file being uploaded (omit to read from STDIN)
    --log      - Logtype of file being uploaded (omit to use base name from file=<path> or 'stream' if STDIN)
    --host     - Hostname or other identifier representing the source of the file being uploaded
    --svcgrp   - Service Group defines a failure domain boundary for anomaly correlation. This allows you to collect logs from multiple
                 applications or support cases and isolate the logs of one from another so as not to mix these
                 in a Root Cause Report. This is referred to as a Service Group in the Zebrium UI.

                 If omitted, Service Group will be set to "default". Default is used to denote a service group that
                 represents shared-services. For example, a database that is shared between two otherwise distinctly separate applications
                 would be considered a shared-service. In this example scenario, you would set the Service Group for one application to "app01"
                 and to "app02" for the other application. For the database logs, you would either omit the --svcgrp setting or you could 
                 explicitly set it do "default" using --svcgrp=default.

                 With this configuration, Root Cause Reports will consider correlated anomalies across:

                     "app01" log events and default (i.e. database logs) and
                     "app02" log events and default (i.e. database logs)

                 but not across:

                     "app01" and "app02"
```

### Advanced Options
Use `ze help_adv` for a complete list of advanced options.

### Batch Uploads
ze supports batch uploads. See [here](https://docs.zebrium.com/docs/setup/ze_batch_uploads/) for usage.

## Examples
1. Ingest three log files associated with the same support case \"sr12345\" (does not assume a .zerc configuration file exists)
```
ze up --file=/casefiles/sr12345/messages.log --svcgrp=sr12345 --host=node01 --log=messages --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>
ze up --file=/casefiles/sr12345/application.log --svcgrp=sr12345 --host=node01 --log=application --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>
ze up --file=/casefiles/sr12345/db.log --svcgrp=sr12345 --host=db01 --log=db --url=<ZE_LOG_COLLECTOR_URL> --auth=<ZE_LOG_COLLECTOR_TOKEN>
```
2. Ingest a continuous tail of /var/log/messages. When reading from a stream (e.g. STDIN) rather than from a file, ze requires the --log flag (assumes a .zerc configuration file exists)
```
tail -f /var/log/messages | ze up --log=varlogmsgs --svcgrp=monitor01 --help=mydbhost
```

## Contributors
* Larry Lancaster (Zebrium)
* Dara Hazeghi (Zebrium)
* Rod Bagg (Zebrium)
* Rob Fair (Zebrium)
