#!/usr/bin/env bash
#set -x
IFS=$'\n' instr=($(cat input.txt))
acc=0
ex=0
unvisited={0..$((${#instr[@]} - 1))}
halt=()
nohalt=()
track=()
array_contains(){
    local arr="$1"
    local seek=$2
    local in=1
    for element in "${arr[@]}"
    do
        [[ "${element}" == "$seek" ]] && in=0 && break
    done
    return $in
}

for ((i=0;i<${#instr[@]};i++))
do
    ((ex++))
    array_contains "${track[@]}" "$i" && nohalt+=("${track[@]}") && break
    track+=($i)
    array_contains "${unvisited[@]}" "$i"  && unset unvisited[$i]
    if [[ ${i} -eq $((${#instr[@]} - 1)) ]] || array_contains "${halt[@]}" $i
    then
        halt+=("${track[@]}")
        break
    elif array_contains "${nohalt[@]}" $i
    then
        nohalt+=("${track[@]}")
        break
    fi
    parts=($(echo ${instr[$i]} | tr " " "n"))
    case ${parts[0]} in
        acc)
            acc=$(($acc ${parts[1]}))
            ;;
        jmp)
            i=$(($i ${parts[1]}))
            ;;
    esac
done
echo $acc



declare -A ran
#i=0
#while [[ -z ${ran[$i]} ]] || [[ $i -eq ${#instr[@]} ]]
#do
#    ran[$i]=$i
#    parts=($(echo ${instr[$i]} | tr " " "\n"))
#    echo -n "$(($i +1)) ${instr[$i]} $(($i+ 1 ${parts[1]})) "
#    case ${parts[0]} in
#        acc) 
#            acc=$(($acc ${parts[1]}))
#            ((i++))
#            ;;
#        jmp)
#            i=$(($i ${parts[1]}))
#            ;;
#        nop)
#            ((i++))
#            ;;
#    esac
#    echo $acc
#done
#echo $acc
