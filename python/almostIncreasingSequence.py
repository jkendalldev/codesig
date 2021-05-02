#!/usr/local/bin/python3

# Given a sequence of integers as an array, determine whether it is possible to
#  obtain a strictly increasing sequence by removing no more than one element from
#  the array.

# For sequence = [1, 3, 2, 1], the output should be false
# almostIncreasingSequence(sequence) = false.
# There is no one element in this array that can be removed in order to
#  get a strictly increasing sequence.

def almostIncreasingSequence(sequence):
  # just count how many times the diff between current number and previous
  # number is <= 0, if that count is > 1, the return false, otherwise return true!
  sequence_did_not_increase_count = 0
  for count, value in enumerate(sequence):
    if count != 0: #if we are not looking at the very first element in the list
      change = value - sequence[count -1]
      if change <= 0:
        sequence_did_not_increase_count += 1
  if sequence_did_not_increase_count > 1:
    return False
  else:
    return True
   

sequence = [1, 5, 3, 1] 
print(almostIncreasingSequence(sequence))


 
 



