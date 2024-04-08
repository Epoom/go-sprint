#!/bin/bash

MOVE_A=no
touch a
touch !
touch "\\"
touch \"
mkdir \'

cp ! \'

if [[ $MOVE_A == "yes" ]]; then
    mv a \'
fi

if [[ $MOVE_A == "no" ]]; then
    rm a
fi