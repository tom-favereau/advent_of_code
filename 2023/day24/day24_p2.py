import numpy as np
from scipy.optimize import linprog
from sympy import symbols, Eq, solve
from scipy.optimize import fsolve

"""
   j'ai d'abord utilisé solve comme ça ne marchait pas j'ai utilisé fsolve, j'arrivais pas à 
    converger. finalement je me suis rendu compte que j'étais très bête d'utiliser 900 équation
    alors que seule 3 sont nécéssaire 
    je suis revenue a solve et j'ai finit
    vue que c'est juste un système asser symble en vérité je devrais pouvoir le transposé 
    en go facilement. j'y reviandrai probablement demain par contre   
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



#print(vel, pos)


n = len(pos)  # Remplacez cela par la valeur réelle de n

# Fonction représentant les équations à résoudre
# Ajouter vx, vy, vz aux variables
vars = np.zeros(n * 3)

minv, maxv = min(v for v in vel[0]), max(v for v in vel[0])
minp, maxp = min(p for p in pos[0]), max(p for p in pos[0])
print(minv, maxv, minp, maxp)

# Initialiser les valeurs initiales pour vx, vy, vz
vars[:3] = np.random.uniform(100000000, 2000000000, 3)
#vars[0] = 274535643646635
#vars[1] = 275608939163845
#vars[2] = 265664851312666
vars[-3:] = np.random.uniform(-10, 10, 3)

#print(vars)



# Fonction représentant les équations à résoudre
def equations_fun(vars):
    x, y, z, vx, vy, vz, T = vars[0], vars[1], vars[2], vars[-3], vars[-2], vars[-1], vars[3:-3]
    equations = []
    for i in range(n):
        equations.append(x + T[i] * vx - pos[i][0] - T[i] * vel[i][0])
        equations.append(y + T[i] * vy - pos[i][1] - T[i] * vel[i][1])
        equations.append(z + T[i] * vz - pos[i][2] - T[i] * vel[i][2])
    return equations

"""
# Résoudre les équations numériquement
#print(len(vars))
solution = fsolve(equations, vars, xtol=1e-8, maxfev=10000)

# Extraire les résultats
x_sol, y_sol, z_sol, vx_sol, vy_sol, vz_sol, T_sol = solution[0], solution[1], solution[2], solution[-3], solution[-2], solution[-1], solution[3:-3]
"""

x0, y0, z0, xv, yv, zv, t1, t2, t3 = symbols('x0 y0 z0 xv yv zv t1 t2 t3')
equations = []
times = [t1, t2, t3]
for i in range(3):
    equations.append(pos[i][0] + vel[i][0]*times[i] - (x0 + xv*times[i]))
    equations.append(pos[i][1] + vel[i][1]*times[i] - (y0 + yv*times[i]))
    equations.append(pos[i][2] + vel[i][2]*times[i] - (z0 + zv*times[i]))

print(equations)

res = solve(equations, x0, y0, z0, xv, yv, zv, t1, t2, t3, dict=True)[0]
print("x: ", res[x0])
print("y: ", res[y0])
print("z: ", res[z0])
print("sum :", res[x0] + res[y0] + res[z0])






file.close()



