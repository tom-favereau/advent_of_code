
import sys


with open(sys.argv[1], 'r') as file:
  ids, ings = file.read().split("\n\n")
  ids, ings = ids.split("\n"), ings.split("\n")
  res1, res2 = 0, 0
  Q = []
  for ing in ings:
    i = int(ing)
    for id in ids:
      a, b = id.split("-")
      a, b = int(a), int(b)
      Q.append((a, b))
      if i >= a and i <= b:
        res1 += 1
        break
  view = []
  while Q:
    a, b = Q.pop()
    new_view = []
    for v1, v2 in view:
        if b < v1 or a > v2:
            new_view.append((v1, v2))
        else:
            a = min(a, v1)
            b = max(b, v2)
    new_view.append((a, b))
    view = new_view
  print(res1, sum(b - a + 1 for a, b in view))
  
