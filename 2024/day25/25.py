import sys

with open(sys.argv[1], 'r') as file: keylocks = file.read().strip().split("\n\n")

keys = []
locks = []
for keylock in keylocks:
    lines = keylock.split("\n")
    rlines = list(reversed(lines))
    if lines[0] == "#####": #is lock
        tmp = []
        for j in range(len(lines[0])):
            for i in range(len(lines)):
                if lines[i][j] == ".":
                    tmp.append(i-1)
                    break
        locks.append(tmp)
    elif rlines[0] == "#####":
        tmp = []
        for j in range(len(rlines[0])):
            for i in range(len(rlines)):
                if rlines[i][j] == ".":
                    tmp.append(i-1)
                    break
        keys.append(tmp)

res1 = 0
for k in keys:
    for l in locks:
        if all(k[i]+l[i] < 6 for i in range(len(k))):
            res1 += 1

print(res1)
