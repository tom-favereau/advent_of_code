import sys
from heapq import heappush, heappop
from collections import deque


with open(sys.argv[1], 'r') as file: maze = list(map(list, file.read().strip().split("\n")))

xs, ys = next((i, j) for i, row in enumerate(maze) for j, val in enumerate(row) if val == "S")





def bfs(xs, ys, maze):


    queue = deque([((xs, ys), 0)])
    visited = set([(xs, ys)])
    D = {}
    D[(xs, ys)] = 0
    dir = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    dir2 = [(0, 2), (2, 0), (0, -2), (-2, 0)]
    res1 = 0
    while queue:
        (x, y), dist = queue.popleft()

        for dx, dy in dir2:
            nx, ny = x + dx, y + dy
            if 0 <= nx < len(maze) and 0 <= ny < len(maze[0]) and (nx, ny) in D :
                if dist - D[(nx, ny)] > 100:
                    res1 += 1


        if maze[x][y] == "E":
            return (D, dist, res1)

        for dx, dy in dir:
            nx, ny = x + dx, y + dy
            if 0 <= nx < len(maze) and 0 <= ny < len(maze[0]) and (nx, ny) not in D and maze[nx][ny] != "#":
                D[(nx, ny)] = dist+1
                queue.append(((nx, ny), dist+1))


    return -1



def solve(D, mc):
    l = [(x, y, D[(x, y)]) for x, y in D]
    l.sort(key=lambda x: x[2])
    res = 0
    for k in range(len(l)):
        for i in range(k):
            x1, y1, d1 = l[k]
            x2, y2, d2 = l[i]
            dist = abs(x2 - x1) + abs(y2 - y1)
            if d1 - d2 > dist and dist <= mc and d1-d2-dist >= 100:
                res += 1
                #break
    return res
D, dist, res1 = bfs(xs, ys, maze) #part 1 plus rapide avec le bfs (linéaire en la taille du chemin au lieu de quadratique)
print(res1)
print(solve(D, 20))
