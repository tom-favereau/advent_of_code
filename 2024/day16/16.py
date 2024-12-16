from collections import defaultdict
import heapq
import sys

with open(sys.argv[1], 'r') as file : mat = list(map(list, file.read().strip().split("\n")))
xs, ys = next((i, j) for i, row in enumerate(mat) for j, val in enumerate(row) if val == "S")
xe, ye = next((i, j) for i, row in enumerate(mat) for j, val in enumerate(row) if val == "E")


def dijkstra(start, end, dir, matrix):
    rows = len(matrix)
    cols = len(matrix[0])
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

    distances = defaultdict(lambda: float('inf'))
    distances[(start[0], start[1], dir)] = 0

    pq = [(0, (start[0], start[1], dir))]
    visited = set()
    parent = {}

    while pq:
        current_dist, current = heapq.heappop(pq)
        i, j, current_dir = current

        if (i, j) == end:
            path = []
            current_node = (i, j, current_dir)
            while current_node in parent:
                path.append(current_node)
                current_node = parent[current_node]
            path.append((start[0], start[1], dir))
            path.reverse()
            return current_dist, (i, j, current_dir), path

        if current in visited:
            continue

        visited.add(current)

        for new_dir in range(4):
            rotation_cost = 0
            if new_dir != current_dir:
                diff = abs(new_dir - current_dir)
                if diff == 3:  # 270° rotation
                    diff = 1
                rotation_cost = 1000 * diff

            di, dj = directions[new_dir]
            new_i, new_j = i + di, j + dj

            if (0 <= new_i < rows and
                    0 <= new_j < cols and
                    matrix[new_i][new_j] != '#'):

                new_cost = current_dist + rotation_cost + 1
                new_node = (new_i, new_j, new_dir)

                if new_cost < distances[new_node]:
                    distances[new_node] = new_cost
                    heapq.heappush(pq, (new_cost, new_node))
                    parent[new_node] = current

    return float('inf'), None, []

cost, end_pos, path = dijkstra((xs, ys), (xe, ye), 0, mat)
print(f"part 1: {cost}")

res2 = []
for i in range(len(mat)):
    for j in range(len(mat[0])):
        if mat[i][j] != "#" and i+j < cost:
            s_to_ij, d, path1 = dijkstra((xs, ys), (i, j), 0, mat)
            _, _, d = d
            ij_to_e, _, path2 = dijkstra((i, j), (xe, ye), d, mat)
            if ij_to_e+s_to_ij <= cost:
                res2.append((i, j, d))

for i, j, d in res2:
    if d == 0:
        mat[i][j] = ">"
    elif d == 1:
        mat[i][j] = "v"
    elif d == 2:
        mat[i][j] = "<"
    else:
        mat[i][j] = "^"

for m in mat:
    print(*m)

print(f"part 2 : {len(res2)}")
