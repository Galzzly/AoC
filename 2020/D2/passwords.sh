#!/usr/bin/env bash
count=0
while read line
do
	vals=($(echo $line |tr " " "\n"))
	min=${vals[0]%-*}
	max=${vals[0]#*-}
	char=${vals[1]%:}
	res=${vals[2]//[^$char]}
	if [[ ${#res} -ge $min && ${#res} -le $max ]]
	then
		((count++))
	fi
done <<< $(cat input)
echo $count
