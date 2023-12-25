import networkx as nx

file = open("input.txt")

lines = file.readlines()

"""
j'ai voulue faire le dernier jour en caml, je code doit être dans AOC_ocaml/bin/day25.ml
mon ford fulkerson ne marche pas bien 

chat gpt m'a conseillé networkx et ça marche bien. 
j'ai fait en go l'algorithme de wikipedia qui est heuristique, ça marche de temps en temps si on laisse tourné longtemps
donc pas concluant 

je suis pas satisfait du tout de mes deux dernier jour, j'ai l'impression d'avoir juste utilisé 
des module python et que le plus gros problème c'etait de lire la doc (et dialoguer avec chatgpt)
"""


indexes = {}
for line in lines:
  s,e = line.split(':')
  for y in e.split():
    try :
        indexes[s].append(y)
    except KeyError:
        indexes[s] = [y]

    try : 
        indexes[y].append(s)
    
    except KeyError:
        indexes[y] = [s]

graph = nx.DiGraph()
for k,vs in indexes.items():
  for v in vs:
    graph.add_edge(k,v,capacity=1)
    graph.add_edge(v,k,capacity=1)

for i in [list(indexes.keys())[0]]:
  for j in indexes.keys():
    if i!=j:
      min_cut, (c1,c2) = nx.minimum_cut(graph, i, j)
      if min_cut == 3:
        print(len(c1)*len(c2))
        break

file.close()
