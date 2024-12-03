import sys
with open(sys.argv[1], "r") as file: l1, l2 = map(sorted, zip(*[(list(map(int, line.split()))) for line in file if line.strip()]))
print(f"{sum(abs(x - y) for x, y in zip(l1, l2))}\n{sum(x * l2.count(x) for x in l1)}")
