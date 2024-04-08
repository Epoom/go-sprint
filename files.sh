#!/bin/bash

touch a
touch '!'
touch "\\"
touch "\""
mkdir "\`"

cp ! "\`"

if [[ $MOVE_A -eq "yes" ]]; then
    mv a "\`"
elif [[ $MOVE_A -eq "no" ]]; then
    rm a
fi