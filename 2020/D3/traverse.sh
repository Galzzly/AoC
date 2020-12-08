#!/usr/bin/env bash
#set -x
input=($(cat input))

# There are 323 entries in this array
# Each entry is 31 characters long

x=0
y=0
count=0

# Don't really need the first entry in the array
while [[ $y -lt ${#input[@]} ]]
do
	((y++))
	((x+=3))
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
