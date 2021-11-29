# ZEBRIUM CLI DETAILS
`ze` is Zebrium's command line interface for uploading log events from files or streams.
## Features

### up (upload)
Upload log event data to your Zebrium instance from a file or stream (stdin) with appropriate meta data.

### help
Display help on ze command usage.

## Getting Started

### Prerequisites

* perl
* perl JSON module
* curl
* Collector token from Zebrium (available from the Log Collector Setup page in the Zebrium UI)
* URL to your instance of Zebrium (available from the Log Collector Setup page in the Zebrium UI)

### Installing

* `git clone https://github.com/zebrium/ze-cli.git`
* move `bin/ze` to appropriate bin directory in your PATH

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

### COMMAND SYNTAX AND OPTIONS
```
  ze up                                                                              \
    [--url=<url>] [--auth=<token>]                                                   \
    [--file=<path>] [--log=<logtype>] [--host=<hostname>] [--svcgrp=<service-group>] \

    --url      - Zebrium Log Collector URL <ZE_LOG_COLLECTOR_URL> (omit to look for url=<url> line in /auto/home/rod/.zerc)
    --auth     - Zebrium Log Collector Token <ZE_LOG_COLLECTOR_TOKEN> (omit to look for auth=<token> line in /auto/home/rod/.zerc)
    --file     - Path to file being uploaded (omit to read from STDIN)
    --log      - Logtype of file being uploaded (omit to use base name from file=<path> or 'stream' if STDIN)
    --host     - Hostname or other identifier representing the source of the file being uploaded
    --svcgrp   - Defines a failure domain boundary for anomaly correlation. This allows you to collect logs from multiple
                 applications or support cases and isolate the logs of one from another so as not to mix these
                 in a Root Cause Report. This is referred to as a Service Group in the Zebrium UI
```

### ADVANCED OPTIONS
Use `ze help` for a complete list of command options.

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
