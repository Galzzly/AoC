#!/usr/bin/python3

import re
from math import prod

def main():
    total = 0

    with open('input.txt', 'r') as input_file:
        lines = input_file.readlines()
        lines = [line.strip() for line in lines]
    
    stars = []
    numbers = {}

    for y in range(len(lines)):
        for number in re.finditer(r'\d+', lines[y]):
            try:
                numbers[y][number.span()] = number.group()
            except KeyError:
                numbers[y] = {number.span(): number.group()}
        for star in re.finditer(r'\*', lines[y]):
            stars.append([y, star.start()])
    
    for star in stars:
        adjacent = []

        for line in range(star[0] - 1, star[0] +2):
            try:
                for number in numbers[line]:
                    if star[1] in range(number[0]-1, number[1]+1):
                        adjacent.append(int(numbers[line][number]))
                        continue
            except KeyError:
                continue
        if len(adjacent) == 2:
            total += prod(adjacent)
    print(total)
        

if __name__ == '__main__':
    main()
