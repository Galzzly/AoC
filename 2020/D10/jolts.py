with open('input') as f:    nums = [int(i) for i in f.readlines()]
target = max(nums) + 3
nums.extend((0, target))
nums.sort()
jolt=0
jolts = [0] * 4
waysTo = {0:1}

for v in nums:
    jolts[v-jolt] += 1
    if v == 0:
        continue
    jolt = v
    waysTo[v] = waysTo.get(v-1, 0) + waysTo.get(v-2, 0) + waysTo.get(v-3, 0) 
print("Part 1: %s" % (jolts[1] * jolts[3]))
print("Part 2: %s" % (waysTo[target]))