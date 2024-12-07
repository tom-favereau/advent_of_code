import sys, functools
with open(sys.argv[1], 'r') as file :lines = map(lambda e : e.split(":"), file.read().strip().split('\n'))
def f(a, s, p2): return a == int(s[0]) if len(s) == 1 else f(a, [str(int(s[0])*int(s[1]))]+s[2:], p2) or f(a, [str(int(s[0])+int(s[1]))]+s[2:], p2) or (p2 and f(a, ["".join(s[0] + s[1])]+s[2:], p2))
print(*functools.reduce(lambda a, e : (a[0]+int(e[0]), a[1]+int(e[0])) if f(int(e[0]), e[1].strip().split(" "), False) else (a[0],a[1]+int(e[0])) if f(int(e[0]), e[1].strip().split(" "), True) else a, lines, (0, 0)))
