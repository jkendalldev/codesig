#!/usr/local/bin/python3

def almostIncreasingSequence(sequence):
  isAlmostIncreasing = True 
  was_not_increased_count = 0

  for count, value in enumerate(sequence):
    if count > 0:
      if value <= sequence[count -1]:
        was_not_increased_count += 1

  if was_not_increased_count > 1:
    isAlmostIncreasing = False
  
  #print(was_not_increased_count)
  return isAlmostIncreasing  

#sequence = [10, 1, 2, 3, 4, 5] 
#sequence = [1, 3, 2]
#sequence = [1, 2, 1, 2]
#sequence = [1, 2, 3, 5, 7, 9, 22]
#sequence = [1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6]
# sequence = [1, 2, 5, 3, 5]
#sequence = [1, 2, 3, 4, 99, 5, 6]
# sequence = [1, 3, 2, 1]

print(almostIncreasingSequence(sequence))


 
 



