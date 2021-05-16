#!/usr/local/bin/python3

# Given an array of strings, return another array containing all of its longest strings.
# For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
# allLongestStrings(inputArray) = ["aba", "vcd", "aba"].
def allLongestStrings(inputArray):
    size = len(max(inputArray, key=len))
    return ([word for word in inputArray if len(word) == size])  #Example of list comprehension

inputArray = ["aba", "aaaaaaa", "aaaaaaaaaaa", "vcd", "zzzzzzzzzzz"]
print(allLongestStrings(inputArray))