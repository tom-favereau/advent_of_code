import sys
from collections import deque

with open(sys.argv[1], 'r') as file : inits, gates = map(lambda e:e.split("\n"), file.read().strip().split("\n\n"))

D = {}
for init in inits:
    var, val = init.split(": ")
    D[var] = int(val)

Q = deque(gates.copy())
while Q:
    gate = Q.popleft()
    g, r = gate.split(" -> ")
    #print(g)
    if "XOR" in g:
        v1, v2 = g.split(" XOR ")
        if v1 in D and v2 in D:
            D[r] = D[v1] ^ D[v2]
        else:
            Q.append(gate)
    elif "OR" in g:
        v1, v2 = g.split(" OR ")
        if v1 in D and v2 in D:
            D[r] = D[v1] or D[v2]
        else:
            Q.append(gate)
    elif "AND" in g:
        v1, v2 = g.split(" AND ")
        if v1 in D and v2 in D:
            D[r] = D[v1] and D[v2]
        else:
            Q.append(gate)
    else:
        print("ERREUR")

res1 = []
for key in D:
    if key[0] == "z":
        res1.append((key, D[key]))

res1.sort()
res1.reverse()
print(int("".join([str(i[1]) for i in res1]), 2))
