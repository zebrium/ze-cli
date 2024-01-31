# Zebrium batch uploads and ze CLI

Zebrium batch uploads provide a way for grouping one or more related uploads so
that they can be monitored and managed later as a unit. Each batch has a unique id
used to identify the batch.

## Batch Uploads vs Service Groups

Batch uploads are different from service groups:

* **Service groups** provide a semantic connection across the data in uploads when looking for incidents.

* **Batch uploads** manage the overall phases of uploading and processing data in
related logs, for example monitoring if a batch is completed, how many lines
of data have been ingested for, the time taken, and so forth.

## Integration into ze cli

Batch uploads are integrated into the `ze` CLI in the following main ways:

* A standalone upload, using the `ze up` CLI, automatically has a batch created for it.
The batch id is displayed when the upload is finished so progress can be
monitored using the  `ze batch state` and `ze batch show` CLIs, described below.

* A set of related uploads, using the `ze up` CLI, can be associated with a specific
batch id that has been created earlier using the `ze batch begin` CLI.
When all the logs for the batch are uploaded, the batch should be completed
using `ze batch end`, or if there are errors the batch can be cancelled
using `ze batch cancel`.
When `ze batch end` is used all the logs for that batch are processed together by Zebrium.

## ze batch CLI subcommand

The `ze batch` CLI subcommand allows batch uploads to be created, completed, cancelled and monitored. It has the syntax:

``` bash
ze batch begin [--url=<url>] [--auth=<auth>] [--batchId=<batchId>]
ze batch end  [--url=<url>] [--auth=<auth>] --batchId=<batchId>
ze batch cancel  [--url=<url>] [--auth=<auth>] --batchId=<batchId>
ze batch state  [--url=<url>] [--auth=<auth>] --batchId=<batchId>
ze batch show  [--url=<url>] [--auth=<auth>] --batchId=<batchId>
```

These options can be access by running `ze batch --help`.  More information can be found by running `ze batch <subcommand> --help`

## Examples

### Upload a large log, monitoring its progress

Upload a log file, on success the new batch id is displayed, usually with a *Processing* state, meaning the log has been accepted by Zebrium and is being scanned for incidents:

```bash
ze up ... --file=myfile.log
State for batch upload baxyz1748ca is Processing
Sent successfully
```

Monitor the batch until processing completes:

```bash
watch ze batch state ... --batchId=baxyz1748ca
```

When the batch upload is completed the state will change, typically to *Done*. For additional information the `ze batch show` option can be used:

```bash
ze batch show ... --batchId=baxyz1748ca

         Batch ID: baxyz1748ca
            State: Done
          Created: 2022-06-08T22:58:18Z
  Completion Time: 2022-06-08T22:59:45Z
  Expiration Time: 2022-06-10T22:59:45Z
            Lines: 377943
  Bundles Created: 2
Bundles Completed: 2
      Upload time: 0:17 min:sec
  Processing time: 1:20 min:sec
```

In this output the expiration time refers to the batch upload metadata, not the uploaded logs or any detected incidents.

### Uploading multiple logs to be processed together

The `ze batch begin` and `ze batch end` subcommands can be used to create a batch upload that spans several linked files. This allows them to be processed as a unit.

Begin a new batch:

```bash
ze batch begin ... 
New batch upload id: baxyz7357473aac1
```

Upload several logs as part of the same batch, using the `--batchId` option:

```bash
ze up --batchId=baxyz7357473aac1 ... --file=file1.log
ze up --batchId=baxyz7357473aac1 ... --file=file2.log
ze up --batchId= baxyz7357473aac1 ... --file=file3.log
```

End the batch:

```bash
ze batch end ... --batchId=baxyz7357473aac1
```

The batch upload can be monitored as in the previous example, using the `ze batch state` and `ze batch show` subcommands.
