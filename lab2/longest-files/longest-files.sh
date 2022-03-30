#!/bin/bash

paths -s -R | grep ".go" | process --project=1 --delimiter=" " | sort -nrk1 | myHead --lines="$1"