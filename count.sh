#!/bin/bash
filecount=$(tree | wc -l)
multiply=$(($filecount * 5))
printf '\t\vTotal files * 5: %s\v\n' "$multiply"