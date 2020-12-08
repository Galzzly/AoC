#!/usr/bin/env bash
declare -A deps rdeps
src='shiny gold'

while read -r line; do
    ctr=${line% bags contain*}
    while read -r n part; do
        deps[$ctr]=${deps[$ctr]-}${deps[$ctr]+','}$part:$n
        rdeps[$part]=${rdeps[$part]-}${rdeps[$part]+','}$ctr
    done < <(sed \
    -e 's/, /\n/g' \
    -e 's/\.$//g' \
    -e 's/ bags\?//g' \
    -e '/^no other$/d' \
    <<< "${line#*bags contain }"
    )
done <<< $(cat input.txt)

stack=("$src")
echo ${#stack[@]}
declare -A seen=([$src]='')
while [[ ${#stack[@]} -ne 0 ]]; do
    i=$((${#stack[@]} - 1))
    echo $i
    part=${stack[$i]}
    unset "stack[$i]"

    if [[ ! -v rdeps[$part] ]]; then continue; fi
    while read -r ctr; do
        if [[ -v seen[$ctr] ]]; then continue; fi
        seen[$ctr]=''
        stack[${#stack[@]}]=$ctr
    done <<< "${rdeps[$part]//,/$'\n'}"
done
echo "$((${#seen[@]} - 1))"

function count {
    local ctr=$1 total=0
    if [[ ! -v deps[$ctr] ]]; then echo 0; return; fi
    while IFS=':' read -r part n; do
        (( total += n * (1 + $(count "$part")) )) || true
    done <<< "${deps[$ctr]//','/$'\n'}"
    echo "$total"
}
count "$src"
