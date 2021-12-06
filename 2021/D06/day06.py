#!/usr/bin/python3

with open('input') as f:
    nums = [int(i) for i in f.readline().strip('\n').split(',')]

count = dict.fromkeys(range(9),0)
for i in range(9):
    for j in range(0, len(nums)):
        if nums[j] == i:
            count[i] += 1

days = 0
while days < 256:
    ncount = dict.fromkeys(range(9),0)
    for i in reversed(range(1,9)):
        ncount[i-1] = count[i]
    if count[0] > 0:
        ncount[6] += count[0]
        ncount[8] += count[0]
    count = ncount
    days += 1
    if days == 80:
        r1 = 0
        for i in range(len(count)):
            r1 += count[i]

        print("Part 1:", r1)

r2 = 0
for i in range(len(count)):
    r2 += count[i]

print("Part 2:", r2)

