#!/usr/bin/env bash

# Whole row range = 0 - 127
# For first seven characters
# F = lower half
# B = upper half
#
# Whole column range = 0 - 7
# For last three characters
# R = upper half
# L = lower half

checkSeat(){
    local e match="$1"
    shift
    for e; do [[ "$e" == "$match" ]] && return 0; done
    return 1
}
res=()

while read line
do
    row=($(seq 0 127))
    col=($(seq 0 7))
    # First 7 chars
    for (( i=0; i<7; i++ ))
    do
        case ${line:$i:1} in
            F) # Lower
                 mid=$((( ${#row[@]} / 2 )))
                 row=( "${row[@]:0:$mid}")
                 ;;
            B) # Upper
                 mid=$((( ${#row[@]} / 2 )))
                 row=( "${row[@]:$mid}" )
                 ;;
        esac
    done
    r=$((( ${row[0]} * 8 )))

    # Last three chars
    for (( i=7; i<${#line}; i++ ))
    do
        case ${line:$i:1} in
            R) # Upper
                mid=$((( ${#col[@]} / 2 )))
                col=( "${col[@]:$mid}" )
                ;;
            L) # Lower
                mid=$((( ${#col[@]} / 2 )))
                col=( "${col[@]:0:$mid}" )
                ;;
            esac
    done
    a=$((( ${col[0]} + $r )))
    res+=( "$a" )
done <<< $(cat input)

max=0
for n in "${res[@]}"
do
    ((n > max)) && max=$n
done

min=$max
for n in "${res[@]}"
do
    ((n < min)) && min=$n
done

echo ${res[@]} | tr " " "\n" | sort -n |awk '$1!=p+1{print p+1"-"$1-1}{p=$1}'
