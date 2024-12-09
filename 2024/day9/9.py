import sys

with open(sys.argv[1], 'r') as file: line = file.read().strip()


l1 = [int(line[i]) for i in range(len(line)) if i%2 == 0]
l2 = [int(line[i]) for i in range(len(line)) if i%2 == 1]
l1.reverse()
l3, l4 = l1.copy(), l2.copy()

i, j, pos, = 0, 0, l1[-1]
res1 = 0
k = 1
while i < len(l1)-k and j < len(l2):
	n = len(l1)-i-1
	if l1[i] < l2[j]:
		l2[j] -= l1[i]
		res1 += n*(2*pos+l1[i]-1)*l1[i]/2
		pos += l1[i]
		i+=1
	
	else:
		res1 += n*(2*pos+l2[j]-1)*l2[j]/2
		pos += l2[j]
		
		if l1[i] == l2[j]:
			i += 1
		else:
			l1[i] -= l2[j]
		res1 += k *(2*pos+l1[-k-1]-1)*l1[-k-1]/2
		pos += l1[-k-1]
		k+= 1
		j+= 1
pos4, pos3 = [l3[-1]], [0]
res2 = 0
for j in range(len(l4)):
	pos4.append(pos4[-1]+l3[-j-2]+l4[j])
for i in range(len(l3)-1):

	pos3.append(pos3[-1]+l4[i]+l3[-i-1])
pos3.reverse()
for i in range(len(l3)):
	n = len(l3)-i-1
	find = False
	for j in range(len(l4)):
		if l4[j] >= l3[i] and pos3[i] > pos4[j]:
			res2 += n *(2*pos4[j]+l3[i]-1)*l3[i]/2
			pos4[j] += l3[i]
			l4[j] -= l3[i]
			find = True

			break
			
	if not find:
		res2 += n *(2*pos3[i]+l3[i]-1)*l3[i]/2

			
			

print(f"{int(res1)}\n{int(res2)}")






		
