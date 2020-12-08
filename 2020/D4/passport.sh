#!/usr/bin/env bash
set -x
IFS= readarray -d '' line < <(awk -v RS= -v ORS='\0' '1' input)

count=0
i=0
while [[ $i -lt ${#line[@]} ]]
do
	parts=($(echo ${line[$i]} | tr " " "\n"))
	if [[ ${#parts[@]} -eq 8 ]]
	then
		((count++))
	elif [[ ${#parts[@]} -eq 7 ]]
	then
		if [[ ! "${parts[@]}" =~ "cid" ]]
		then
			#continue
		#else
			((count++))		
		fi
	fi
	((i++))
	echo $count
done 
echo $count
