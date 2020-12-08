#!/usr/bin/env bash

IFS=$'\n' instr=($(cat input.txt))
tochange=()
for ((i=0;i<${#instr[@]};i++))
do
    parts=($(echo ${instr[$i]} |tr " " "\n"))
    [[ "${parts[0]}" != "acc" ]] && tochange+=($i)
done
acc=0
attempt=0
for j in "${tochange[@]}"
do
    ((attempt++))
    echo "Attempt: $attempt"
    declare -A ran
    i=0
    while [[ -z ${ran[$i]} ]]
    do
        ran[$i]=$i
        parts=($(echo ${instr[$i]} | tr " " "\n"))
        case ${parts[0]} in
            acc)
                acc=$(($acc ${parts[1]}))
                ((i++))
                ;;
            jmp)
                if [[ $i -eq $j ]]
                then
                    # act as NOP
                    ((i++))
                else
                    i=$(($i ${parts[1]}))
                fi
                ;;
            nop)
                if [[ $i -eq $j ]]
                then
                    # ant as JMP
                    i=$(($i ${parts[1]}))
                else
                    ((i++))
                fi
                ;;
        esac
    done
    [[ $i -ge $((${#instr[@]} - 1)) ]] && break
    unset ran
    acc=0
done
echo $acc in $attempt
