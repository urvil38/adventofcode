import sys
from collections import Counter

infile = sys.argv[1] if len(sys.argv) >= 2 else 'inputs/day02.in'

D = open(infile).read().strip()

reports=[]
for line in D.split('\n'):
  reports.append(list(map(int, line.split())))


def is_safe(r):
  isSafe = True
  inc_or_dec = (r==sorted(r) or r==sorted(r, reverse=True))
  for i in range(len(r)-1):
    diff = abs(r[i]-r[i+1])
    if not 1 <= diff <= 3:
      isSafe = False
  return isSafe and inc_or_dec


def part1(reports):
  part1=0
  for r in reports:
      part1 += is_safe(r)
  print(part1)

def part2(reports):
  part2=0
  for r in reports:
    isSafe=False
    for i in range(len(r)):
      r1 = r[:i] + r[i+1:]
      if is_safe(r1):
        isSafe=True
    if isSafe:
      part2+=1
  print(part2)

part1(reports)
part2(reports)