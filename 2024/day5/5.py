import sys, functools
with open(sys.argv[1], 'r') as file : order, lines = file.read().strip().split('\n\n')
lines, order, s = list(map(lambda l: list(map(int, l.split(','))), lines.split('\n'))), list(map(lambda l: list(map(int, l.split('|'))), order.split('\n'))), lambda l : sorted(l, key=functools.cmp_to_key(lambda a, b: -1 if [a, b] in order else (1 if [b, a] in order else 0))) 
print(*functools.reduce(lambda acc, x : [acc[0] + s(x)[len(x)//2], acc[1]] if x == s(x) else [acc[0], acc[1] + s(x)[len(x)//2]], lines, [0, 0]))
