#!/bin/zsh

set -euo pipefail

# check if user gave an argument
if [[ "$#" -ne 1 ]]; then
  echo "One argument required!"
  exit 1
fi

# variables
sum=0
counter=0

for size in $(
  paths -R -s | grep ".$1" | process --project=1 --delimiter=" "
)
do
  echo "Added new file"
  sum=$((sum + size))
  counter=$((counter + 1))
done

# print average file size
if [[ "$counter" -ne 0 ]]; then
  echo "$((sum / counter))"
else
  echo "There were no files with extension $1!"
fi