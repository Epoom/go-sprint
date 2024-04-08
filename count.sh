#!/bin/bash
filecount=$(find . | wc -l)
multiply=$(($filecount * 5))
printf '\t\vTotal files * 5: %s\v\n' "$multiply"