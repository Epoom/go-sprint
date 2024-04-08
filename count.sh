#!/bin/bash
filecount=$(ls -R | wc -l)
multiply=$(($filecount * 5))
printf '\t\v Total files * 5: %s\v\n' "$multiply"