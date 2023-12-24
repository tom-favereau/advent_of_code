import numpy as np

"""
    j'avais commençais à l'écrire en go mais quand j'ai écris la matrice 
    je me suis dit que j'aurais mieux fait de le faire avec np.linalg 
"""

file = open("input.txt")


lines = file.readlines()

pos = []
vel = []
for i in lines:
    sep = i.split(" @ ")
    p = sep[0].split(", ")
    pos.append((int(p[0]), int(p[1]), int(p[2])))
    v = sep[1].split(", ")
    vel.append((int(v[0]), int(v[1]), int(v[2])))



res = 0
for i in range(len(pos)):
    for j in range(len(pos)):
        if i != j:
            vi = vel[i]
            vj = vel[j]
            pi = pos[i]
            pj = pos[j]

            A = np.array([[-vi[0], vj[0]], [-vi[1], vj[1]]])
            Y = np.array([pi[0]-pj[0], pi[1]-pj[1]])
            try:
                X = np.linalg.solve(A, Y)
                x = pi[0]+X[0]*vi[0]
                y = pi[1]+X[0]*vi[1]

                if 200000000000000<=x<=400000000000000 and 200000000000000<=y<=400000000000000 and X[0] >= 0 and X[1] >= 0:
                    res += 1
                    #print("solv", x, y, X[0])

            except :
                pass
           
print(res/2)

file.close()
