import sys

with open(sys.argv[1], 'r') as file : lines = file.read().strip().split("\n")


def f(n):
    res = n
    res = (res ^ (res*64))%16777216
    res = (res ^ res // 32)%16777216
    res = (res ^ (res*2048))%16777216
    return res



res1 = 0
D = {}
for line in lines:
    x = int(line)
    m = int(str(x)[-1])
    pc = []
    for _ in range(2000):
        x = f(x)
        pc.append((int(str(x)[-1])-m, int(str(x)[-1])))
        m = int(str(x)[-1])
    D2 = {}
    for i in range(3, len(pc)):
        a, b, c, d = pc[i-3][0], pc[i-2][0], pc[i-1][0], pc[i][0]
        p = pc[i][1]
        if (a, b, c, d) not in D2:
            D2[(a, b, c, d)] = p

    for key in D2:
        if key in D:
            D[key] += D2[key]
        else:
            D[key] = D2[key]

    res1 += x


print(res1)
print(D[max(D, key=D.get)])
