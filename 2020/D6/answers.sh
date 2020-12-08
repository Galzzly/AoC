#!/usr/bin/env bash

#IFS= readarray -d '' line < <(awk -v RS= -v ORS='\0' '1' input)
count=0
while read line
do
    numP=($(echo ${line} | tr " " "\n"))
    numAns=$(grep -o . <<< $(echo ${line}) |sort -u |wc -l)
    [[ ${#numP[@]} -gt 1 ]] && ((numAns--))
    count=$(($count + $numAns))
done <<< $( perl -00 -ne 's/\n/ /g;print "$_\n";' input )
echo $count


