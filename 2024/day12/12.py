import sys
from collections import deque

with open(sys.argv[1], 'r') as file: mat = file.read().strip().split('\n')
mat = [list(row) for row in zip(*mat[::-1])]
n = len(mat)
m=len(mat[0])


v = set()
def bfs(a, b, e):
	q = deque([(a, b)])
	vis = set()
	vis.add((a, b))
	v.add((a, b))
	fr=set()
	ans1, ans2 = 0, 0
	while q:
		i, j = q.popleft()
		v.add((i, j))
		vis.add((i, j))
		t=0
		if i<n-1 and mat[i+1][j] == e and (i+1, j) not in vis:
			q.append((i+1, j))
			v.add((i+1, j))
			vis.add((i+1, j))
		if i>0 and mat[i-1][j] == e and (i-1, j) not in vis:
			q.append((i-1, j))
			v.add((i-1, j))
			vis.add((i-1, j))
		if j<m-1 and mat[i][j+1] == e and (i, j+1) not in vis:
			q.append((i, j+1))
			v.add((i, j+1))
			vis.add((i, j+1))
		if j>0 and mat[i][j-1] == e and (i, j-1) not in vis:
			q.append((i, j-1))
			v.add((i, j-1))
			vis.add((i, j-1))

		if not (i<n-1 and mat[i+1][j] == e):
		        ans1+=1
		        fr.add(("i", i+1, j,i))
		if not (i>0 and mat[i-1][j] == e):
		        ans1+=1
		        fr.add(("i", i-1, j,i))
		if not (j<m-1 and mat[i][j+1] == e):
		        ans1+=1
		        fr.add(("j", i, j+1, j))
		if not (j>0 and mat[i][j-1] == e):                
		        ans1+=1
		        fr.add(("j", i, j-1, j))

		t=ans2
		if not (i<n-1 and mat[i+1][j] == e) and ("i", i+1, j+1, i) not in fr and ("i", i+1, j-1,i) not in fr:
		        ans2+=1
		if not( i>0 and mat[i-1][j] == e) and ("i", i-1, j+1,i) not in fr and ("i", i-1, j-1, i) not in fr:
		        ans2+=1
		if not (j<m-1 and mat[i][j+1] == e) and ("j", i+1, j+1,j) not in fr and ("j", i-1, j+1, j) not in fr:
		        ans2+=1
		if not (j>0 and mat[i][j-1] == e) and ("j", i+1, j-1, j) not in fr and ("j", i-1, j-1, j) not in fr:                
		        ans2+=1


	return ans1, ans2

res1, res2 = 0, 0
for i in range(n):
        for j in range(m):
                if (i,j) not in v:
                        fr=set()
                        tmp=len(v)
                        p, s =bfs(i,j,mat[i][j])
                        a = len(v)-tmp
                        res1+=a*p
                        res2+=a*s
                        
print(res1)   
print(res2)      
                        
