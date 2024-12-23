import sys

with open(sys.argv[1], 'r') as file: lines = file.read().strip().split("\n")

G = {}
for line in lines:
    a, b = line.split("-")
    if a in G:
        G[a].append(b)
    else:
        G[a] = [b]
    if b in G:
        G[b].append(a)
    else:
        G[b] = [a]

def triangles(G):
    triangles = []
    sommets = list(G.keys())

    for i in range(len(sommets)):
        for j in range(i + 1, len(sommets)):
            for k in range(j + 1, len(sommets)):
                sommet1 = sommets[i]
                sommet2 = sommets[j]
                sommet3 = sommets[k]

                if (sommet2 in G[sommet1] and
                        sommet3 in G[sommet1] and
                        sommet3 in G[sommet2]):
                    triangles.append([sommet1, sommet2, sommet3])

    return triangles


def clique_max(G):
    def bron_kerbosch(R, P, X, cliques_maximales):
        if not P and not X:  
            if len(R) > len(cliques_maximales[0]):
                cliques_maximales[0] = R.copy()
            return
        pivot = max((len(set(G[v]) & P) for v in P | X), default=0)
        for v in list(P):
            voisins = set(G[v])
            bron_kerbosch(R | {v}, P & voisins, X & voisins, cliques_maximales)
            P.remove(v)
            X.add(v)

    R = set()  
    P = set(G.keys())  
    X = set()  
    cliques_maximales = [set()] 
    bron_kerbosch(R, P, X, cliques_maximales)

    return list(cliques_maximales[0])

lt = triangles(G)
res1 = 0
for t in lt:
    if t[0][0] == "t" or t[1][0] == "t" or t[2][0] == "t":
        res1 += 1


print(res1)
print(",".join(sorted(clique_max(G))))
