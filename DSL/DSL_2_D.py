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

    def update(self, s, t, x):
        data = self._data
        counter = self._counter
        l = s + self._offset
        r = t + 1 + self._offset
        while l < r:
            if l & 1 == 0:
                data[l] = (counter, x)
            if r & 1 == 0:
                data[r - 1] = (counter, x)
            l = l // 2
            r = (r - 1) // 2
        self._counter += 1

    def find(self, i):
        data = self._data
        j = i + self._offset
        t = data[j]
        while j >= 1:
            j = (j - 1) // 2
            t = max(t, data[j])
        return t[1]


readline = stdin.readline

n, q = map(int, readline().split())

st = RangeUpdateSegmentTree(n)
st.update(0, n - 1, 2 ** 31 - 1)
result = []
for _ in range(q):
    query = readline()
    if query[0] == '0':
        _, s, t, x = map(int, query.split())
        st.update(s, t, x)
    elif query[0] == '1':
        _, i = map(int, query.split())
        result.append(st.find(i))
print(*result, sep='\n')
