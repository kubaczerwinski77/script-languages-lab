#!/bin/bash

set -euo pipefail

# check if user passes exactly 3 args
if [[ "$#" -ne 3 ]]; then
  echo "You have to pass 3 args in order: \"sum-cases [continent] [year] [month]\""
  exit 1
fi

cat covid.tsv.txt |
  process --delimiter=$'\t' --separator=";" --project=10,3,2,4 | # pick columns continent | year | month | cases
  grep "$1;$2;$3" | # grab only those which fits arguments passsed to function
  process --delimiter=";" --project=3 | # pick only column with cases
  aggregate --using=sum # aggreggate values using sum