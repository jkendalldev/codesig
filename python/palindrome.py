#!/usr/local/bin/python3

def checkPalindrome(forwardString):
  reverseString = forwardString[::-1]
  if (forwardString == reverseString):
    return True
  else:
    return False

checkPalindrome("noon")


 
 



