#!/bin/bash


function usage() {
    echo "batch_upload [options] file1...."
    echo "Options:"
    echo "  -h              - Print this help message"
    echo "  -u <url>        - URL for upload, if not using .zerc file"
    echo "  -a <auth        - Authentication token, if not using .zerc file"
    echo "  -o <ze_options> - Other Ze options"
    echo ""
    echo "Example, specifying service group and host:"
    echo "    batch_upload  -o '--svcgrp=ABC --host=myhost' *.log"
}

ZE_ARGS=""
ZE_ARGS2=""
while getopts "o:u:a:h" options; do         # Loop: Get the next option;
    case "${options}" in
    h) usage
       exit 1;;
    u)
      ZE_ARGS="$ZE_ARGS --url='$OPTARG'";;
    a)
      ZE_ARGS="$ZE_ARGS --auth='$OPTARG'";;
    o)
      ZE_ARGS2="$ZE_ARGS2 $OPTARG";;
    ?)
      echo "Invalid argument '$1'"
      exit 1;;
    *)
      break;;
  esac
done

shift $(($OPTIND - 1))
if [ $# -eq 0 ]; then
    echo "No files specified"
    exit 1
fi
for f in "$@"
do
    if [ ! -r "$f" ]; then
        echo "Log $f cannot be read"
        exit 1
    fi
done

BATCH_ID=`eval ze batch begin $ZE_ARGS | awk '{print $5}'`
if [ -z "$BATCH_ID" ]; then
    echo "Unable to create batch upload"
    exit 1
fi
echo "Started batch upload $BATCH_ID"
for f in "$@"
do
   echo "Uploading $f..."
   eval  "ze up --batch_id=$BATCH_ID $ZE_ARGS $ZE_ARGS2 --file='$f'"
done
eval "ze batch end $ZE_ARGS --batch_id=$BATCH_ID"
echo "Use 'ze batch show --batch_id=$BATCH_ID $ZE_ARGS' to monitor batch progress"

