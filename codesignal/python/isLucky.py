#!/usr/local/bin/python3

# Ticket numbers usually consist of an even number of digits. 
# A ticket number is considered lucky if the sum of the first half of the digits 
# is equal to the sum of the second half.

# Given a ticket number n, determine if it's lucky or not.
# Example
# For n = 1230, the output should be
# isLucky(n) = true;
# For n = 239017, the output should be
# isLucky(n) = false.

def isLucky(n):
    numList = [int(digit) for digit in str(n)]
    half = len(numList)//2
    return(sum(numList[:half]) == sum(numList[half:]))


n = 1230
print(isLucky(n))