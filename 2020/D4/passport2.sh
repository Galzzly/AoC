#!/usr/bin/env bash
IFS= readarray -d '' line < <(awk -v RS= -v ORS='\0' '1' input)

count=0
i=0
while [[ $i -lt ${#line[@]} ]]
do
	parts=($(echo ${line[$i]} | tr " " "\n"))
	if [[ ${#parts[@]} -eq 8 ]] || [[ ${#parts[@]} -eq 7 && ! "${parts[@]}" =~ "cid" ]]
	then
		invalid=0
		for part in ${parts[@]}
		do
			case ${part} in 
				byr*)
					val=${part#*:}
					if [[ ${#val} -eq 4 ]] && [[ ${val} -ge 1920 ]] && [[ ${val} -le 2002 ]]
					then
						continue
					else
						((invalid++))
					fi
					;;
				iyr*)
					val=${part#*:}
					if [[ ${#val} -eq 4 ]] && [[ ${val} -ge 2010 ]] && [[ ${val} -le 2020 ]]
					then
						continue
					else
						((invalid++))
					fi
					;;
				eyr*)
					val=${part#*:}
					if [[ ${#val} -eq 4 ]] && [[ ${val} -ge 2020 ]] && [[ ${val} -le 2030 ]]
					then
						continue
					else
						((invalid++))
					fi
					;;
				hgt*)
					val=${part#*:}
					case ${val} in
						*cm)
							num=${val%cm}
							if [[ ${num} -ge 150 ]] && [[ ${num} -le 193 ]]
							then
								continue
							else
								((invalid++))
							fi
							;;
						*in)
							num=${val%in}
							if [[ ${num} -ge 59 ]] && [[ ${num} -le 76 ]]
							then
								continue
							else
								((invalid++))
							fi
							;;
						*)	((invalid++))
							;;
					esac
					;;
				hcl*)
					val=${part#*:}
					if [[ ${val} =~ ^#([0-9]|[a-f]){6}$ ]]
					then
						continue
					else
						((invalid++))
					fi
					;;
				ecl*)
					val=${part#*:}
					case ${val} in
						amb|blu|brn|gry|grn|hzl|oth) continue ;;
						*) ((invalid++)) ;;
					esac
					;;
				pid*)
					val=${part#*:}
					if [[ ${val} =~ ^[0-9]{9}$ ]]
					then
						continue
					else
						((invalid++))
					fi
					;;
			esac
		done
		if [[ ${invalid} -eq 0 ]]
		then
			((count++))
		fi
	fi
	((i++))
done 
echo $count
