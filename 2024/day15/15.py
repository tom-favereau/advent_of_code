import sys
with open(sys.argv[1], 'r') as file: mat, ins = file.read().strip().split("\n\n")
mat, ins, x, y = list(map(list, mat.split("\n"))), "".join(ins.split("\n")), 0, 0
x, y = next((i, j) for i, row in enumerate(mat) for j, val in enumerate(row) if val == "@")


mmat = []
for row in mat:
    nr = []
    for i in row:
        if i == '@':
            nr.extend(['@', '.'])
        elif i == 'O':
            nr.extend(['[', ']'])
        elif i == '#':
            nr.extend(['#', '#'])
        else:
            nr.extend(['.', '.'])
    mmat.append(nr)

###Part1####

for k in range(len(ins)):
  a = ins[k]
  if a == "^":
    if mat[x-1][y] == "#":
      pass
    elif mat[x-1][y] == "O":
      tmpi, tmpj = x-1, y
      while mat[tmpi][tmpj] == "O":
        tmpi -= 1
      if mat[tmpi][tmpj] == "#":
        pass
      else:
        mat[x-1][y] = "@"
        mat[tmpi][tmpj] = "O"
        mat[x][y]  = "."
        x -= 1
    else:
      mat[x-1][y] = "@"
      mat[x][y] = "."
      x -= 1
  elif a == ">":
    if mat[x][y+1] == "#":
      pass
    elif mat[x][y+1] == "O":
      tmpi, tmpj = x, y+1
      while mat[tmpi][tmpj] == "O":
        tmpj += 1
      if mat[tmpi][tmpj] == "#":
        pass
      else:
        mat[x][y+1] = "@"
        mat[tmpi][tmpj] = "O"
        mat[x][y]  = "."
        y += 1
    else:
      mat[x][y+1] = "@"
      mat[x][y] = "."
      y += 1
  elif a == "<":
    if mat[x][y-1] == "#":
      pass
    elif mat[x][y-1] == "O":
      tmpi, tmpj = x, y-1
      while mat[tmpi][tmpj] == "O":
        tmpj -= 1
      if mat[tmpi][tmpj] == "#":
        pass
      else:
        mat[x][y-1] = "@"
        mat[tmpi][tmpj] = "O"
        mat[x][y]  = "."
        y -= 1
    else:
      mat[x][y-1] = "@"
      mat[x][y] = "."
      y -= 1
  elif "v":
    if mat[x+1][y] == "#":
      pass
    elif mat[x+1][y] == "O":
      tmpi, tmpj = x+1, y
      while mat[tmpi][tmpj] == "O":
        tmpi += 1
      if mat[tmpi][tmpj] == "#":
        pass
      else:
        mat[x+1][y] = "@"
        mat[tmpi][tmpj] = "O"
        mat[x][y]  = "."
        x += 1
    else:
      mat[x+1][y] = "@"
      mat[x][y] = "."
      x += 1

res = 0
for i in range(len(mat)):
  for j in range(len(mat[0])):
    if mat[i][j] == "O":
      res += 100*i+j
res1 = res


###Part2###

x, y = next((i, j) for i, row in enumerate(mmat) for j, val in enumerate(row) if val == "@")




x, y = 0, 0
for i in range(len(mmat)):
  for j in range(len(mmat[0])):
    if mmat[i][j] == "@":
      x, y = i, j
      break

D = {}
def f(i, j, m, dir):
  if (i, j, dir) in D:
    return D[(i, j, dir)]
  if dir == "^":
    tmp = m[i-1][j]
    if tmp == "#":
      D[(i, j, dir)] = False
      return False
    elif tmp == "[":
      res = f(i-1, j, m, dir) and f(i-1, j+1, m, dir)
      D[(i, j, dir)] = res
      return res 
    elif tmp == "]":
      res = f(i-1, j, m, dir) and f(i-1, j-1, m, dir)
      D[(i, j, dir)] = res
      return res
    else:
      D[(i, j, dir)] = True
      return True
  elif dir == ">":
    tmp = m[i][j+1]
    if tmp == "#":
      D[(i, j, dir)] = False
      return False
    elif tmp == "[":
      res = f(i, j+1, m, dir)
      D[(i, j, dir)] = res
      return res 
    elif tmp == "]":
      res = f(i, j+1, m, dir) 
      D[(i, j, dir)] = res
      return res
    else:
      D[(i, j, dir)] = True
      return True
  elif dir == "<":
    tmp = m[i][j-1]
    if tmp == "#":
      D[(i, j, dir)] = False
      return False
    elif tmp == "[":
      res = f(i, j-1, m, dir)
      D[(i, j, dir)] = res
      return res 
    elif tmp == "]":
      res = f(i, j-1, m, dir)
      D[(i, j, dir)] = res
      return res
    else:
      D[(i, j, dir)] = True
      return True
  elif dir == "v":
    tmp = m[i+1][j]
    if tmp == "#":
      D[(i, j, dir)] = False
      return False
    elif tmp == "[":
      res = f(i+1, j, m, dir) and f(i+1, j+1, m, dir)
      D[(i, j, dir)] = res
      return res 
    elif tmp == "]":
      res = f(i+1, j, m, dir) and f(i+1, j-1, m, dir)
      D[(i, j, dir)] = res
      return res
    else:
      D[(i, j, dir)] = True
      return True
  
def g(i, j, m, dir):
  if (i, j, dir) in D:
    return
  else:
    D[(i, j, dir)] = True
  if dir == "^":
    tmp = m[i-1][j]
    if tmp == "#":
      return
    elif tmp == "[":
      g(i-1, j, m, dir)
      g(i-1, j+1, m, dir)
      m[i-1][j] = m[i][j]
      m[i][j] = "."
    elif tmp == "]":
      g(i-1, j-1, m, dir)
      g(i-1, j, m, dir)
      m[i-1][j] = m[i][j]
      m[i][j] = "."
    else:
      m[i-1][j] = m[i][j]
      m[i][j] = "."

  if dir == ">":
    tmp = m[i][j+1]
    #print("ok g", tmp)
    if tmp == "#":
      return
    elif tmp == "[":
      g(i, j+1, m, dir)
      m[i][j+1] = m[i][j]
      m[i][j] = "."
    elif tmp == "]":
      g(i, j+1, m, dir)
      m[i][j+1] = m[i][j]
      m[i][j] = "."
    else:
      m[i][j+1] = m[i][j]
      m[i][j] = "."
  if dir == "<":
    tmp = m[i][j-1]
    if tmp == "#":
      return
    elif tmp == "[":
      g(i, j-1, m, dir)
      m[i][j-1] = m[i][j]
      m[i][j] = "."
    elif tmp == "]":
      g(i, j-1, m, dir)
      m[i][j-1] = m[i][j]
      m[i][j] = "."
    else:
      m[i][j-1] = m[i][j]
      m[i][j] = "."
  elif dir == "v":
    tmp = m[i+1][j]
    if tmp == "#":
      return
    elif tmp == "[":
      g(i+1, j, m, dir)
      g(i+1, j+1, m, dir)
      m[i+1][j] = m[i][j]
      m[i][j] = "."
    elif tmp == "]":
      g(i+1, j-1, m, dir)
      g(i+1, j, m, dir)
      m[i+1][j] = m[i][j]
      m[i][j] = "."
    else:
      m[i+1][j] = m[i][j]
      m[i][j] = "."
      return


for k in range(len(ins)):
  a = ins[k]
  D = {}
  #print()
  #print(a)
  #for m in mmat:
  #  print(*m)
  if f(x, y, mmat, a):
    #print("ok")
    D = {}
    g(x, y, mmat, a)
    if a == "^":
      x -= 1
    elif a == ">":
      y += 1
    elif a == "<":
      y -= 1
    else:
      x += 1
  else:
    pass

res = 0
for i in range(len(mmat)):
  for j in range(len(mmat[0])):
    if mmat[i][j] == "[":
      res += 100*i+j

res2 = res
print(res1)
print(res2)
