import sys
from heapq import heappush, heappop

with open(sys.argv[1], 'r') as file: coords = [tuple(map(int, e.split(','))) for e in file.read().strip().split("\n")]

def dijkstra(matrix):
    if not matrix or not matrix[0]:
        return -1

    n, m = len(matrix), len(matrix[0])

    distances = [[float('inf')] * m for _ in range(n)]
    distances[0][0] = 0

    pq = [(0, 0, 0)]

    directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    while pq:
        dist, x, y = heappop(pq)

        if x == n - 1 and y == m - 1:
            return dist

        if dist > distances[x][y]:
            continue

        for dx, dy in directions:
            new_x, new_y = x + dx, y + dy

            if (0 <= new_x < n and
                    0 <= new_y < m and
                    matrix[new_x][new_y] == '.'):

                new_dist = dist + 1

                if new_dist < distances[new_x][new_y]:
                    distances[new_x][new_y] = new_dist
                    heappush(pq, (new_dist, new_x, new_y))

    return -1


def solve(n):
    m = []
    for i in range(71):
        tmp = []
        for j in range(71):
            if (j, i) in coords:
                tmp.append(".")
            else:
                tmp.append(".")
        m.append(tmp)

    for i in range(n):
        a, b = coords[i]
        m[b][a] = "#"

    return dijkstra(m)

def dichot():
    left = 0
    right = len(coords)

    if solve(right) >= 0:
        return None

    if solve(left) < 0:
        return left

    while left + 1 < right:
        mid = (left + right) // 2

        if solve(mid) >= 0:
            left = mid
        else:
            right = mid

    return right-1

print(solve(1024))
print(coords[dichot()])
