#!/bin/bash


for ((i=1; i<=1000; i++))
do
    compareOutput=$(bash linuxcomparer.sh)

echo -e "$compareOutput"
if [[ "$compareOutput" == *"FAILED"* ]]; then
    break
fi
done
