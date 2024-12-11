import sys
with open(sys.argv[1], 'r') as file: lines=list(map(lambda e:list(map(int, e)), file.read().strip().split('\n')))

s = set()
def f(i,j, pr):
    if lines[i][j]==9:
        s.add((i,j))
        return 1
    tmp=0
    if i<len(lines)-1 and lines[i+1][j] == pr+1:
        tmp += f(i+1, j, pr+1)
    if j<len(lines[0])-1 and lines[i][j+1] == pr+1:
        tmp += f(i, j+1, pr+1)
    if i>0 and lines[i-1][j] == pr+1:
        tmp += f(i-1, j, pr+1)
    if j>0 and lines[i][j-1] == pr+1:
        tmp += f(i, j-1, pr+1)
    return tmp
                           
res1 = 0
res2 = 0
for i in range(len(lines)):
    for j in range(len(lines[0])):
        if lines[i][j] == 0:
            s = set()
            tmp=f(i,j, 0)
            res1+=len(s)
            res2+=tmp

print(res1, res2)
