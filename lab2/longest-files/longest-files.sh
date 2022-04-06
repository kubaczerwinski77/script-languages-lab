#!/bin/bash

set -euo pipefail

# check if user gave an argument
if [[ "$#" -ne 1 ]]; then
  echo "One argument required (number of lines)!"
  exit 1
fi

paths -s -R | grep ".go" | process --project=1 --delimiter=" " | sort -nrk1 | myHead --lines="$1"