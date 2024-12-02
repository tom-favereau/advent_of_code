import sys
with open(sys.argv[1]) as f:
    lines = [list(map(int, line.split())) for line in f if line.strip()]
    is_safe = lambda s: all(s[i] - s[i + 1] in {1, 2, 3} for i in range(len(s) - 1)) or all(s[i] - s[i + 1] in {-1, -2, -3} for i in range(len(s) - 1))
    print(f"{sum(is_safe(l) for l in lines)}\n{sum(any(is_safe(l[:k] + l[k+1:]) for k in range(len(l))) for l in lines)}")
