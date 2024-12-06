import sys
from collections import Counter

infile = sys.argv[1] if len(sys.argv) >= 2 else 'inputs/day01.in'

D = open(infile).read().strip()

LL=[]
RL=[]
C = Counter()
for line in D.split('\n'):
  L,R = line.split()
  L,R = int(L), int(R)
  LL.append(L)
  RL.append(R)
  C[R]+=1

part1=0
for a, b in zip(sorted(LL), sorted(RL)):
  part1+=abs(a-b)
print(part1)

part2=0
for a in LL:
  part2 += (a * C[a])
print(part2)