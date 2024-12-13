import re, sys
with open(sys.argv[1], 'r') as file:blocks, res1, res2 =file.read().strip().split('\n\n'), 0, 0
for b in blocks:
        ba, bb, p = re.search(r"Button A: X\+(\d+), Y\+(\d+)", b), re.search(r"Button B: X\+(\d+), Y\+(\d+)", b), re.search(r"Prize: X=(\d+), Y=(\d+)", b)
        (x1, y1), (x2, y2), (X, Y) = map(int, ba.groups()), map(int, bb.groups()), map(int, p.groups())   
        a, b, c, d = (y2*X-x2*Y)/(x1*y2-y1*x2), (x1*Y-y1*X)/(x1*y2-y1*x2), (y2*(X+10000000000000)-x2*(Y+10000000000000))/(x1*y2-y1*x2), (x1*(Y+10000000000000)-y1*(X+10000000000000))/(x1*y2-y1*x2)
        res1, res2 = (res1+3*a+b if int(a)==a and int(b)==b else res1, res2+3*c+d if int(c)==c and int(d)==d else res2)
print(f"{int(res1)}\n{int(res2)}")
