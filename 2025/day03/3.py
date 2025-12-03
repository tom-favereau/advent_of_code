import sys
with open(sys.argv[1], 'r') as file:
  lines = file.read().split("\n")
  res1, res2 = 0, 0
  for l in lines:
    if l == "":
      break
    ll = list(map(int, list(l)))

    m1, m2 = ll[0], ll[1]
    m = [(ll[i], i) for i in range(12)]
    for i in range(1, len(ll)):
      for j in range(12):
        if ll[i] > m[j][0] and (j==0 or m[j-1][1]<i) and i+(11-j) < len(l):
          for k in range(12-j):
            m[j+k] = (ll[i+k], i+k) 
          break
      if ll[i] > m1 and i+1 != len(l):
        m1 = ll[i] 
        m2 = ll[i+1]
      elif ll[i] > m2:
        m2 = ll[i]
    res1 += 10*m1+m2
    res2 += sum(m[k][0]*10**(11-k) for k in range(12))
    #print(m1, m2, 10*m1+m2)
  print(res1, res2)
  print
      