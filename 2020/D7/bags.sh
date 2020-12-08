#!/usr/bin/env bash
set -x
# Need to check the bags that can hold gold bags directly
# Then check the bags that hold the bags that can hold gold bags
declare -A deps rdeps
while read line
do
    clr=${line% bags contain*}
    while read n part
    do
        deps[$clr]=$part:$n
        rdeps[$part]=${rdeps[$part]-}${rdeps[$part]+','}$clr
    done < <(sed \
        -e 's/, /\n/g' \
        -e 's/\.$//g' \
        -e 's/ bags\?//g' \
        -e '/^no other$/d' \
        <<< "${line#*bags contain}"
    )
done <<< $(cat input.txt)

src='shiny gold'
stack=("$src")
declare -A seen=([$src]='')
while [[ ${#stack[@]} -ne 0 ]]
do
    i=$((${#stack[@]} - 1))
    part=${stack[$i]}
    unset "stack[$i]"

    [[ ! -v rdeps[$part] ]] && continue
    while read clr
    do
        [[ -v seen[$clr] ]] && continue
        seen[$clr]=''
        stack[${#stack[@]}]=$clr
    done <<< "${rdeps[$part]//,/$'\n'}"
done
echo "$((${#seen[@]} - 1))"
