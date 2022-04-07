#!/bin/bash

set -euo pipefail

# check if user passes exactly 3 args
if [[ "$#" -ne 3 ]]; then
  echo "You have to pass 3 args in order: \"sum-cases [continent] [year] [month]\""
  exit 1
fi

# validate if year is not-negative number
if [[ "$2" -lt 0 ]]; then
  echo "Year cannot be a negative number! Given: $2"
  exit 1
fi

# validate if month is a number between 1 and 12
if [[ "$3" -le 1 || "$3" -ge 12 ]]; then
  echo "Month must be a number between 1 and 12! Given: $3"
  exit 1
fi

cat covid.tsv |
  process --delimiter=$'\t' --separator=";" --project=10,3,2,4 | # pick columns continent | year | month | cases
  grep "$1;$2;$3" | # grab only those which fits arguments passsed to function
  process --delimiter=";" --project=3 | # pick only column with cases
  aggregate --using=sum --separator="" # aggreggate values using sum