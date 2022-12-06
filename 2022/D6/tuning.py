import time

start = time.time()
f = "input.txt"
line = open(f).read().strip()

def streamcheck(length: int) -> int:
    for i in range(len(line)):
        mark = line[i:i+length]
        if len(set(mark)) == length:
            return i+length
        
print(f"Part 1: {streamcheck(4)}")
print(f"Part 2: {streamcheck(14)}")

print("Total Time: " + str(time.time() - start))