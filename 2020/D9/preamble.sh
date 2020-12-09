#!/usr/bin/env bash
set -x
nums=($(cat input))

found=''
for (( i=25; i<${#nums[@]}; i++ ))
do
    notFound=false
    echo "Checking ${nums[$i]}..."
    for (( j=$((${i} - 25 )); j<$i; j++ ))
    do
        [[ ${nums[$j]} -lt ${nums[$i]} ]] && continue
        target=$((${nums[$i]} - ${nums[$j]}))
        for (( k=$((${i} - 25 )); k<$i; k++ ))
        do
            [[ $k -ne $j ]] && continue
            case ${nums[$k]} in
                ${target})  notFound=true
                            break 
                            ;;
            esac
        done
    done
    if $notFound 
    then
        found=${nums[$i]}
        break
    fi
done

echo $found

