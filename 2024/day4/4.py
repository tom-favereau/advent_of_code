import sys
with open(sys.argv[1], 'r') as file: lines = file.read().strip().split("\n")
cols, d = list(map(lambda e : "".join(e), list(zip(*lines)))), list(map(lambda e : "".join(e), [[lines[i][j] for i in range(len(lines)) for j in range(len(lines[i])) if i - j == k] for k in range(-len(lines) + 1, len(lines))] + [[lines[i][j] for i in range(len(lines)) for j in range(len(lines[i])) if i + j == k]              for k in range(len(lines) + len(lines) - 2, -1, -1)]))
count = lambda lines : sum(l.count("XMAS") + "".join(reversed(list(l))).count("XMAS") for l in lines)
print(count(lines)+count(cols)+count(d), sum((i+2 < len(lines) and j+2 < len(lines[0]) and lines[i][j] + lines[i+1][j+1] + lines[i+2][j+2] in {"MAS", "SAM"} and lines[i][j+2] + lines[i+1][j+1] + lines[i+2][j] in {"MAS", "SAM"}) for i in range(len(lines)) for j in range(len(lines[0]))))
