#!/bin/bash

loopcount=$1

if  [[ $loopcount -gt 100 ]] ; then
    number=100
else 
    number=$loopcount
fi

for ((i=1; i<=$number; i++)) ; do
    echo "$i times I've printed egertyakopoom"
done