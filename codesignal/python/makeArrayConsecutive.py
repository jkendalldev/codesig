#!/usr/local/bin/python3

def makeArrayConsecutive2(statues):
  sorted_statues = sorted(statues)
  missing_statues = 0
  increase = 0
  for count, value in enumerate(sorted_statues):
    if count != 0:
      increase = value - sorted_statues[count -1]
      if increase != 1:
        missing_statues = missing_statues + (increase - 1)
  return missing_statues   

  
statues = [6, 2, 3, 8]
print(makeArrayConsecutive2(statues))


 
 



