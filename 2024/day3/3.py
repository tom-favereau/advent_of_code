import sys, re
with open(sys.argv[1], 'r') as file: line, p, p2 = "do()" + file.read() + "don't()", r"mul\((\d{1,3}),(\d{1,3})\)", r"do\(\).*?don't\(\)"
f = lambda l : sum(int(a)*int(b) for a, b in re.findall(p, l))
print(f"{f(line)}\n{sum(f(l) for l in re.findall(p2, line, re.DOTALL))}")
