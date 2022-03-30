#!/bin/bash

# set default properties
src="."
dst=

# get flag values (--src, --dst)
for i in "$@"; do
  case $i in
    --src=*)
      src="${i#*=}"
      shift # past argument=value
      ;;
    --dst=*)
      dst="${i#*=}"
      shift # past argument=value
      ;;
    -*|--*)
      echo "Unknown option $i"
      exit 1
      ;;
  esac
done

# check if given src is a directory
if [[ ! -d "$src" ]]; then
    echo "Given --src ($src) is not a directory!"
    usage
fi

# check if given dst is not an empty string
if [[ -z "$dst" ]]; then
    echo "No destination given, destination is required!"
    usage
fi

cp -R "$src" "$dst"