# ze-cli

Zebrium's command line interface for uploading log events from files or streams, viewing log events and the definitions of event types in the database.

## Features
##### upload
Upload log event data to your Zebrium database instance from a file or stream (stdin) with appropriate meta data.
##### def
Show the event-type (eType) definition for structured events in the database.
##### cat
Show events from the database by: meta-data, eType, time range, or first occurence in CSV, JSON, pretty-print or raw format.
## Getting Started
##### Prerequisites
* Perl
* cURL
* API token from Zebrium
* URL to your instance of Zebrium
##### Installing
* git clone https://github.com/zebrium/ze-cli.git
* move bin/ze to appropriate bin directory in your PATH
## Configuration
No preconfiguration is required. All options can be specified as command line argumants. However, see **Setup** section below for information on configuring your .zerc file.
##### Setup

For convenience, your API token and your Zebrium instance URL can be specified in the .zerc file located in your home directory.
**Format of .zerc file**
```
auth=YOUR_API_TOKEN
url=https://YOUR_ZE_INSTANCE_NAME.zebrium.com
```
##### Environment Variables

## Testing your installation

## Contributors

## Contributing
