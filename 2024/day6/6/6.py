import sys
with open(sys.argv[1], 'r') as file: lines = list(map(list, file.read().strip().split("\n")))
(x, x2, y, y2), res1, res2, d, direc, stop1 = next(((i, i, j, j) for i in range(len(lines)) for j in range(len(lines[0])) if lines[i][j] == "^"), (None, None, None, None)), set(), 0, 0, ["U", "R", "D", "L"], False
for (i, j) in [(i, j) for i in range(len(lines)) for j in range(len(lines[0]))]: 
		(lines[i][j], stop1, save, d, x, y, seen) = ("#", False, lines[i][j], 0, x2, y2, set()) if lines[i][j] == "." else (lines[i][j], True, lines[i][j], 0, x2, y2, set())
		while not stop1:
			stop2 = False
			while not stop2:
				tmpx, tmpy, (x, y) = x, y, (x-1, y) if direc[d%4] == "U" else (x, y+1) if direc[d%4] == "R" else (x+1, y) if direc[d%4] == "D" else (x, y-1)	
				(stop1, stop2, res2, d, x, y, _) = (True, True, res2, d, x, y, print(len(res1)) if (i, j) == (0, 0) else None) if x >= len(lines) or x < 0 or y >= len(lines[0]) or y < 0 else (True, True, res2+1, d+1, tmpx, tmpy, None) if (lines[x][y] == "#" and (d%4, x, y) in seen) else (stop1, True, res2, d+1, tmpx, tmpy, seen.add((d%4, x, y))) if lines[x][y] == "#" else (stop1, stop2, res2, d, x, y, res1.add((x, y)))
				
				
		lines[i][j] = save
print(res2)
