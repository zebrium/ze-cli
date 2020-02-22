# ZEBRIUM CLI DETAILS
# ze-cli
Zebrium's command line interface for uploading log events from files or streams, viewing log events and the definitions of event types in the database.
## Features
##### upload
Upload log event data to your Zebrium instance from a file or stream (stdin) with appropriate meta data.
##### def
Show the event-type (eType) definition for structured events in the database.
<!--
##### cat
Show events from the database by: meta-data, eType, time range, or first occurrence in CSV, JSON, pretty-print or raw format.
-->
## Getting Started
##### Prerequisites
* perl
* perl JSON module
* curl
* API auth token from Zebrium
* URL to your instance of Zebrium
##### Installing
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
No configuration is required. All options can be specified as command line arguments. However, see **Setup** section below for information on configuring your .zerc file.
##### Setup
For convenience, your API token and your Zebrium instance URL can be specified in your $HOME/.zerc file.

Example .zerc file:
```
auth=YOUR_ZE_API_AUTH_TOKEN
url=https://YOUR_ZE_API_INSTANCE_NAME.zebrium.com
```
##### Environment Variables
None
## Usage
Use `ze help` for a complete list of command operations and options.
```
ze help
```
## Examples
1. Ingest the log file /var/log/messages (does not assume a .zerc configuration file exists)
```
ze up --file=/var/log/messages --ids='node_name=node01,node_id=234fd3e1-2a34' --auth=YOUR_AUTH_TOKEN --url=https://YOUR_ZE_API_INSTANCE_NAME.zebrium.com
```
2. Ingest a continuous tail of /var/log/messages. When reading from a stream (e.g. STDIN) rather than from a file, ze requires the --log flag (assumes a .zerc configuration file exists)
```
tail -f /var/log/messages | ze up --log=varlogmsgs --ids='node_name=node01,node_id=234fd3e1-2a34'
```
3. Show 20 events (using pretty-printed JSON) already ingested into your Zebrium instance (assumes a .zerc configuration file exists)
```
ze cat --lim=20 --fmt=pp
```
## Contributors
* Larry Lancaster (Zebrium)
* Dara Hazeghi (Zebrium)
* Rod Bagg (Zebrium)
