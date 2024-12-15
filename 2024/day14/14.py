import sys
import numpy as np


with open(sys.argv[1], 'r') as file: lines = file.read().strip().split('\n')
n, m = 101, 103
tl, bl, tr, br = 0, 0, 0, 0

k = 0
continuer = True
while continuer:
  mat = [[0 for _ in range(m)] for _ in range(n)]
  for line in lines:
    a, b = line.split(" ")
    a = a[2:]
    b = b[2:]
    px, py = map(int, a.split(","))
    vx, vy = map(int, b.split(","))
    pfx = (px+vx*k)%n
    pfy = (py+vy*k)%m
    if k== 100 and pfx > n//2 and pfy > m//2:
      br += 1
    elif k==100 and pfx > n//2 and pfy < m//2:
      tr += 1
    elif k==100 and pfx < n//2 and pfy > m//2:
      bl += 1
    elif k==100 and pfx < n//2 and pfy < m//2:
      tl += 1
    mat[pfx][pfy] += 1
  test = True
  for i in range(len(mat)):
    for j in range(len(mat[1])):
      if mat[i][j] > 1:
        test = False
  if test:
    print(k)
    #for l in mat:
    #  print(l)
    continuer = False
    break
  k+=1


  #print(np.array(mat))

print(br, bl, tr, tl)
print(br*bl*tr*tl)  
print(k)
