import sys
with open(sys.argv[1], 'r') as file: lines = list(map(int, file.read().strip().split(' ')))

d = {i:1 for i in lines}
for k in range(75):
    if k == 25:
        print(sum(d[i] for i in d))
    dd = {}
    for i in d:
        tmp=str(i)
        if d[i] == 0:
            pass
        elif i == 0 :
            dd[1] = dd[1]+d[0] if 1 in dd else d[0]
        elif len(tmp)%2==0:
            a=int(tmp[len(tmp)//2:])
            b=int(tmp[:len(tmp)//2])
            dd[b] = dd[b]+d[i] if b in dd else d[i]
            dd[a] = dd[a]+d[i] if a in dd else d[i]
        else:
            dd[2024*i] = dd[2024*i]+d[i] if 2024*i in dd else d[i]
                    
    d=dd

print(sum(d[i] for i in d))
