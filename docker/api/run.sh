#!/bin/bash

make run_api_server

cat <<EOF >>~/.bashrc
trap 'pkill -TERM golang; sleep 3;exit 0' TERM
EOF
exec /bin/bash
