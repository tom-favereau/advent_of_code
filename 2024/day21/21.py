from itertools import permutations
import sys

with open(sys.argv[1], 'r') as f:
    codes = f.read().strip().split("\n")

digit = {
    '7': (0, 0), '8': (0, 1), '9': (0, 2),
    '4': (1, 0), '5': (1, 1), '6': (1, 2),
    '1': (2, 0), '2': (2, 1), '3': (2, 2),
    '0': (3, 1), 'A': (3, 2)
}

dir = {
    '^': (0, 1), 'A': (0, 2),
    '<': (1, 0), 'v': (1, 1), '>': (1, 2)
}

memo = {}


def f(rid, cur, dst, n):
    if (rid, cur, dst, n) in memo:
        return memo[(rid, cur, dst, n)]

    keys = digit if rid == 0 else dir
    cp = keys[cur]
    tp = keys[dst]  # tp pour target position au lieu de dp
    dx = tp[0] - cp[0]
    dy = tp[1] - cp[1]

    if rid == n - 1:
        return abs(dx) + abs(dy) + 1

    s = []
    for _ in range(abs(dx)):
        s.append('^' if dx < 0 else 'v')
    for _ in range(abs(dy)):
        s.append('<' if dy < 0 else '>')

    if not s:
        return 1

    c = []
    for r in permutations(s):
        pos = cp
        tot = 0
        valid = True
        for i, k in enumerate(r):
            tot += f(rid + 1, 'A' if i == 0 else r[i - 1], k, n)

            if k == '^':
                pos = (pos[0] - 1, pos[1])
            elif k == 'v':
                pos = (pos[0] + 1, pos[1])
            elif k == '<':
                pos = (pos[0], pos[1] - 1)
            else:
                pos = (pos[0], pos[1] + 1)

            if not any(v == pos for v in keys.values()):
                valid = False
                break

        if valid:
            tot += f(rid + 1, r[-1], 'A', n)
            c.append(tot)

    res = min(c)
    memo[(rid, cur, dst, n)] = res
    return res


res1 = res2 = 0
for c in codes:
    l1 = f(0, 'A', c[0], 3)
    l2 = f(0, 'A', c[0], 26)
    for i in range(1, len(c)):
        l1 += f(0, c[i - 1], c[i], 3)
        l2 += f(0, c[i - 1], c[i], 26)
    n = int(c[:-1])
    res1 += l1 * n
    res2 += l2 * n

print(res1)
print(res2)
