S = int(input())

h = S // 3600
m = (S % 3600) // 60
s = (S % 3600) % 60
print('%d:%d:%d' % (h, m, s))
