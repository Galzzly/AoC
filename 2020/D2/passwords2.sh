#!/usr/bin/env bash
set -x
count=0
while read line
do
	vals=($(echo $line |tr " " "\n"))
	p1=${vals[0]%-*}
	p2=${vals[0]#*-}
	((p1--))
	((p2--))
	char=${vals[1]%:}
	r1=${vals[2]:$p1:1}
	r2=${vals[2]:$p2:1}
	if [[ "${r1}" == "${char}" && "${r2}" == "${char}" ]]
	then
		continue
	elif [[ "${r1}" == "${char}" || "${r2}" == "${char}" ]]
	then
		((count++))
	fi
done <<< $(cat input)
echo $count
