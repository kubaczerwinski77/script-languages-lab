#!/bin/bash

set -euo pipefail

# check if user passes exactly 2 args
if [[ "$#" -ne 2 ]]; then
  echo "You have to pass 2 args in order: \"worst-months [numberOfMonths] [country]\""
  exit 1
fi

# validate if month is a number between 1 and 12
if [[ "$1" -le 1 || "$1" -ge 12 ]]; then
  echo "Month must be a number between 1 and 12! Given: $1"
  exit 1
fi

cat covid.tsv.txt |
  process --delimiter=$'\t' --separator=";" --project=6,3,2 --select="$2" | # pick country, year, month that matches arg
  sort -t ';' -nk2 -nk3 | # sort by year and then by month
  uniq | # save unique groups to file
  while read group 
  do
    cat covid.tsv.txt | # read file
    process --delimiter=$'\t' --separator=";" --project=6,3,2,5 | # pick country, year, month, deaths
    grep "$group" | # choose only rows that matches unique group
    process --delimiter=";" --project=3 | # pick only deaths column
    aggregate --using=avg --label="$group" # calc average and add label
  done |
  sort -t ';' -nrk4 | # sort by the avg for all the groups
  myHead --lines="$1" # print first N number of lines
