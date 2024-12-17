import sys
import time

with open(sys.argv[1], 'r') as file :reg, prog = file.read().strip().split("\n\n")

"""
2,4 B = A%8
1,3 B = B^3 => B = 3
7,5 C = A// 2**B => C = A // 8
4,2 B = B^C 
0,3 A = A// 2**3
1,5 B = B^5
5,5 out(B)
3,0 if A == 0 go 0
"""

reg = reg.split("\n")
a = reg[0].split(" ")
a = int(a[-1])
b = reg[1].split(" ")
b = int(b[-1])
c = reg[2].split(" ")
c = int(c[-1])

prog = prog.split(" ")
prog = list(map(int, prog[1].split(",")))


def f(a, b, c):
    A, B, C = [a], [b], [c]
    combo = [[0], [1], [2], [3], A, B, C]
    i = 0
    res = []
    while i < len(prog):
        ins, op = prog[i], prog[i+1]
        #print(ins, op)
        if ins == 0: #adv
            A[0] = int(A[0]/(2**combo[op][0]))
        elif ins == 1: #bxl
            B[0] = B[0] ^ op
        elif ins == 2: #bst
            B[0] = combo[op][0]%8
        elif ins == 3: # jnz
            if A[0] != 0:
                i = op
                continue
        elif ins == 4: # bxc
            B[0] = B[0] ^ C[0]
        elif ins == 5: #out
            #print(f"{combo[op][0]%8},", end="")
            res.append(str(combo[op][0]%8))
        elif ins == 6: #bdv
            B[0] = int(A[0] / (2 ** combo[op][0]))
        elif ins == 7: # cdv
            C[0] = int(A[0] / (2 ** combo[op][0]))
        i += 2
    return res

print(",".join(list(map(str, f(a, b, c)))))

k = 16
i = 0
a2 = 0
target = list(map(str, prog))
res = []
A = [0]
while k > 0:
    nA = []
    for pra2 in A:
        for i in range(8):
            a2 = pra2 + 8**(k-1)*i
            res = f(a2, b, c)
            if len(target) == len(res) and res[k-1] == target[k-1]:
                nA.append(a2)

    A = nA
    k -= 1

print(A)
