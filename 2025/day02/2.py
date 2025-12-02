import sys
with open(sys.argv[1], 'r') as file: 
  ranges = file.read().split(",")
  res1 = 0
  res2 = 0
  for r in ranges:
    a, b = r.split("-")
    a, b = int(a), int(b)
    for i in range(a, b+1):
      si = str(i)
      n = len(si)
      if n%2==0 and si[:(n//2)] == si[(n//2):]:
        res1+=i
        res2+=i
      elif n%3==0 and si[:(n//3)] == si[(n//3):2*(n//3)] and si[:(n//3)] == si[2*(n//3):]:
        res2+=i 
      elif n%5 == 0 and all(si[i*(n//5):(i+1)*(n//5)] == si[:(n//5)] for i in range(5)):
        res2+=i
      elif n%7 == 0 and all(si[i*(n//7):(i+1)*(n//7)] == si[:(n//7)] for i in range(7)):
        res2+=i
  print(res1, res2)
