#!/bin/bash

number=$1

if  (( number > 100 )) ; then
    number=100
fi

for (( a=1; a<=number; a++))
do
    echo "$number times I've printed egertyakopoom"
done