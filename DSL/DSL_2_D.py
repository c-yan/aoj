from sys import stdin


class RangeUpdateSegmentTree:
    def __init__(self, size):
        self._size = size
        t = 1
        while t < size:
            t *= 2
        self._offset = t - 1
        self._data = [(-1, 0)] * (t * 2 - 1)
        self._counter = 0

    def range_update(self, start, stop, x):
        data = self._data
        counter = self._counter
        l = start + self._offset
        r = stop + self._offset
        while l < r:
            if l & 1 == 0:
                data[l] = (counter, x)
            if r & 1 == 0:
                data[r - 1] = (counter, x)
            l = l // 2
            r = (r - 1) // 2
        self._counter += 1

    def __getitem__(self, key):
        data = self._data
        i = key + self._offset
        t = data[i]
        while i > 0:
            i = (i - 1) // 2
            t = max(t, data[i])
        return t[1]


readline = stdin.readline

n, q = map(int, readline().split())

st = RangeUpdateSegmentTree(n)
st.range_update(0, n, 2 ** 31 - 1)
result = []
for _ in range(q):
    query = readline()
    if query[0] == '0':
        _, s, t, x = map(int, query.split())
        st.range_update(s, t + 1, x)
    elif query[0] == '1':
        _, i = map(int, query.split())
        result.append(st[i])
print(*result, sep='\n')
