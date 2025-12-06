import sys
with open(sys.argv[1], 'r') as file:
  lines = file.read().split("\n")
  del lines[-1]
  m = []
  res1, res2 = 0, 0
  for line in lines:
    if line == "":
      continue
    tmp = list(filter(lambda x: x!=" " and x!="", line.split(" ")))
    m.append(tmp)
  for i in range(len(m[0])):
    if m[-1][i] == "+":
      tmp = 0
      for j in range(len(m)-1):
        tmp += int(m[j][i])
      res1 += tmp
    else:
      tmp = 1
      for j in range(len(m)-1):
        tmp *= int(m[j][i])
      res1 += tmp

  l = []
  for i in range(len(lines[0])):
    k = len(lines[0]) - i - 1
    tmp = "".join([lines[j][k] for j in range(len(lines)-1)])    
    if tmp.strip() != "":
      l.append(int(tmp))
    if lines[-1][k] == "+":
      res2 += sum(l)
      l = []
    elif lines[-1][k] == "*":
      tmp = 1
      for j in l:
        tmp *= j
      res2 += tmp
      l = []
      
  print(res1, res2)