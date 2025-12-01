
import sys
with open(sys.argv[1], 'r') as file:
  lines = file.read().split("\n")
  cnt = 50
  res1=0
  res2 = 0
  for l in lines:
    if l == "":
      pass
    elif l[0] == "R":
      for i in range(int(l[1:])):
        cnt += 1
        cnt=cnt % 100
        if cnt == 0:
          res2+=1
    else:
      for i in range(int(l[1:])):
        cnt -= 1
        cnt = cnt%100
            
        if cnt == 0:
          res2+=1
    if cnt == 0:
      res1 += 1
    
print(res1, res2)



