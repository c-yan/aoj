from sys import stdin


class RangeAddSegmentTree:
    def __init__(self, size):
        self._size = size
        t = 1
        while t < size:
            t *= 2
        self._offset = t - 1
        self._data = [0] * (t * 2 - 1)

    def range_add(self, start, stop, x):
        data = self._data
        l = start + self._offset
        r = stop + self._offset
        while l < r:
            if l & 1 == 0:
                data[l] += x
            if r & 1 == 0:
                data[r - 1] += x
            l = l // 2
            r = (r - 1) // 2

    def __getitem__(self, key):
        data = self._data
        i = key + self._offset
        result = data[i]
        while i > 0:
            i = (i - 1) // 2
            result += data[i]
        return result


readline = stdin.readline

n, q = map(int, readline().split())

st = RangeAddSegmentTree(n)
result = []
for _ in range(q):
    query = readline()
    if query[0] == '0':
        _, s, t, x = map(int, query.split())
        st.range_add(s - 1, t, x)
    elif query[0] == '1':
        _, t = map(int, query.split())
        result.append(st[t - 1])
print(*result, sep='\n')
