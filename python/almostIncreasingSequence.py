#!/usr/local/bin/python3

def almostIncreasingSequence(sequence):
  isAlmostIncreasing = True
  threw_out_a_number_count = 0
  
  for count, value in enumerate(sequence):
    if count > 0:
      previous_num = sequence[count -1]
      if value <= previous_num:
        threw_out_a_number_count += 1
        # only do this if we are NOT on a count of 1...
        if count != 1 and count + 1 < (len(sequence)):
          next_num = sequence[count +1]
          if next_num <= previous_num:
            isAlmostIncreasing = False
           
  if threw_out_a_number_count > 1:
    isAlmostIncreasing = False
  
  return isAlmostIncreasing  

#sequence = [10, 1, 2, 3, 4, 5] 
#sequence = [1, 3, 2]
#sequence = [1, 2, 1, 2]
#sequence = [1, 2, 3, 5, 7, 9, 22]
#sequence = [1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6]
sequence = [1, 2, 5, 3, 5]

print(almostIncreasingSequence(sequence))


 
 



