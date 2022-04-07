#!/bin/bash

set -euo pipefail

# check if user passed 2 arguments
if [[ "$#" -ne 2 ]]; then
  echo "You have to pass 2 args in order: \"ratio [continent] [year]\""
  exit 1
fi

CASES=$(cat covid.tsv |
  process --delimiter=$'\t' --separator=";" --project=10,3,4 |
  grep "$1;$2" |
  process --delimiter=";" --project=2 |
  aggregate --using=sum --separator="")
CASES=${CASES//".00"/""}

DEATHS=$(cat covid.tsv |
  process --delimiter=$'\t' --separator=";" --project=10,3,5 |
  grep "$1;$2" |
  process --delimiter=";" --project=2 |
  aggregate --using=sum --separator="")
DEATHS=${DEATHS//".00"/""}

if [[ "$CASES" -eq 0 ]]; then
  echo "There is not enought data to calculate death/cases ratio!"
  exit 1
fi

printf "Death rate in %s in %s was %.5f%%\n" $1 $2 $((10**6 * DEATHS/CASES))e-6
