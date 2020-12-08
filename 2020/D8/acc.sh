#!/usr/bin/env bash
#set -x
IFS=$'\n' instr=($(cat input.txt))
acc=0
array_contains(){
    local arr="$1"
    local seek=$2
    local in=1
    for element in "${arr[@]}"
    do
        [[ $element == "$seek" ]] && in=0 && break
    done
    return $in
}
declare -A ran
i=0
while [[ -z ${ran[$i]} ]] || [[ $i -eq ${#instr[@]} ]]
do
    ran[$i]=$i
    parts=($(echo ${instr[$i]} | tr " " "\n"))
    echo -n "$(($i +1)) ${instr[$i]} $(($i+ 1 ${parts[1]})) "
    case ${parts[0]} in
        acc) 
            acc=$(($acc ${parts[1]}))
            ((i++))
            ;;
        jmp)
            i=$(($i ${parts[1]}))
            ;;
        nop)
            ((i++))
            ;;
    esac
    echo $acc
done
echo $acc
