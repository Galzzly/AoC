#!/usr/bin/env bash
#set -x
input=($(cat input))

# There are 323 entries in this array
# Each entry is 31 characters long

x=0
y=0
count=0

r=$1
d=$2

# Don't really need the first entry in the array
while [[ $y -lt ${#input[@]} ]]
do
	((y+=$d))
	((x+=$r))
	if [[ $x -ge ${#input[$y]} ]]
	then
		((x -= ${#input[$y]}))
	fi
	res=${input[$y]:$x:1}
	if [[ "$res" == "#" ]]
	then
		((count++))
	fi
done
echo $count
