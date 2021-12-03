file = open('input', 'r')
lines = file.readlines()
p1 = [0,0]
p2 = [0,0,0]

for line in lines:
  s = line.split()
  v = int(s[1])
  match s[0]:
    case 'forward':
        p1[0] += v
        p2[0] += v
        p2[1] += v*p2[2]
    case 'down':
        p1[1] -= v
        p2[2] += v
    case 'up':
        p1[1] += v
        p2[2] -= v

print("Part 1:", (abs(p1[0]) * abs(p1[1])), "Part 2:", (abs(p2[0]) * abs(p2[1])))