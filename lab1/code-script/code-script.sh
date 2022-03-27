#!/bin/bash

# PATH TO FILE = /Users/kubaczerwinski77/Desktop/studies/sem-let-2021-22/skryptowe/skryptowe-lab/lab1/return-code/test.txt

# Validate script argument count
if test "$#" -eq 0; then
    echo "At least 1 argument is required!"
    exit 1
fi

# Validate if input file exists
if test ! -f "$1"; then
    echo "File with this path does not exist!"
    exit 1
fi

# Get script arguments
filename="$1"
shift
keywords="$@"

# Display params to user
echo "Scanning file $(basename $filename) for these words: $keywords"

# Run the executable
return-code < "$filename" $keywords 
exit_code="$?"

# Display output to user
if test "$exit_code" -gt 0; then
    echo "The word passed as the $exit_code argument was the most common."
else
    echo "There was no such keywords in this text!"
fi