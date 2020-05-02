#!/usr/bin/env bash

if [[ "$1" = "--ts" ]]; then
    psql -x postgres://db_admin:pass@0.0.0.0:35433/zlab
elif [[ "$1" = "--pg" ]]; then
    psql -x postgres://db_admin:pass@0.0.0.0:35434/zlab
else
    echo "need to provide target"
    exit 1
fi
