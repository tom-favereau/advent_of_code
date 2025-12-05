import sys
with open(sys.argv[1], 'r') as file:
  lines = file.read().split("\n")
  mat = [list(l) for l in lines]
  n, m = len(mat), len(mat[0])
  res1, res2 = 0, 0
  for i in range(n):
    for j in range(m):
      if mat[i][j] != "@":
        continue
      cnt = 0
      if i > 0 and j > 0 and mat[i-1][j-1] == "@":
        cnt += 1
      if i > 0 and j < m-1 and mat[i-1][j+1] == "@":
        cnt += 1

      if i < n-1 and j > 0 and mat[i+1][j-1] == "@":
        cnt += 1
      if i < n-1 and j < m-1 and mat[i+1][j+1] == "@":
        cnt += 1

      if i > 0 and mat[i-1][j] == "@":
        cnt += 1
      if j > 0 and mat[i][j-1] == "@":
        cnt += 1
        
      if i < n-1 and mat[i+1][j] == "@":
        cnt += 1
      if j < m-1 and mat[i][j+1] == "@":
        cnt += 1
      if cnt < 4:
        res1 += 1
  print(res1)

  while True:
    tmp = 0
    for i in range(n):
      for j in range(m):
        if mat[i][j] != "@":
          continue
        cnt = 0
        if i > 0 and j > 0 and mat[i-1][j-1] == "@":
          cnt += 1
        if i > 0 and j < m-1 and mat[i-1][j+1] == "@":
          cnt += 1

        if i < n-1 and j > 0 and mat[i+1][j-1] == "@":
          cnt += 1
        if i < n-1 and j < m-1 and mat[i+1][j+1] == "@":
          cnt += 1

        if i > 0 and mat[i-1][j] == "@":
          cnt += 1
        if j > 0 and mat[i][j-1] == "@":
          cnt += 1
        
        if i < n-1 and mat[i+1][j] == "@":
          cnt += 1
        if j < m-1 and mat[i][j+1] == "@":
          cnt += 1
        if cnt < 4:
          tmp += 1
          mat[i][j] = "x"
    if tmp == 0:
      break
    res2 += tmp
print(res2)
