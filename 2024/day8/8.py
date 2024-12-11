import sys


file = open(sys.argv[1])
lines = file.read().strip().split("\n")

d = {}
res1 = set()
res2 = set()
test = set()
for i in range(len(lines)):
  for j in range(len(lines[0])):
    if lines[i][j] == "#":
      test.add((i, j))
    elif lines[i][j] != ".":
      if lines[i][j] in d.keys():
        for a, b in d[lines[i][j]]:
          #p1 = ((min(i, a)-abs(a-i)), min(j, b)-abs(b-j))
          #p2 = ((max(i, a)+abs(a-i)), (max(j,b)+abs(b-j)))
          p1 = (i-(a-i), j-(b-j))
          p2 = (a-(i-a), b-(j-b))
          #print(a, i, b, j, p1, p2)
          if p1[0] >= 0 and p1[0] < len(lines) and p1[1] >= 0 and p1[1] < len(lines[0]):
            res1.add(p1)
          if p2[0] >= 0 and p2[0] < len(lines) and p2[1] >= 0 and p2[1] < len(lines[0]):
            res1.add(p2)
          while p1[0] >= 0 and p1[0] < len(lines) and p1[1] >= 0 and p1[1] < len(lines[0]):
            res2.add(p1)
            #print("p1", a-i, b-j, p1)
            p1 = (p1[0]-(a-i), p1[1]-(b-j))
          while p2[0] >= 0 and p2[0] < len(lines) and p2[1] >= 0 and p2[1] < len(lines[0]):
            res2.add(p2)
            #print("p2",i-a, j-b, p2)
            p2 = (p2[0]-(i-a), p2[1]-(j-b))
        d[lines[i][j]].append((i, j))
      else:
        d[lines[i][j]] = [(i, j)]
for k in d.keys():
  for p in d[k]:
    if len(d[k]) != 1:
      res2.add(p)
print(len(res1))
print(len(res2))

file.close()
