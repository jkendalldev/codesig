#!/usr/local/bin/python3

# Write a function that reverses characters in (possibly nested) parentheses 
# in the input string.

# Input strings will always be well-formed with matching ()s.

# Example

# For inputString = "(bar)", the output should be
# reverseInParentheses(inputString) = "rab";

# For inputString = "foo(bar)baz", the output should be
# reverseInParentheses(inputString) = "foorabbaz";

# For inputString = "foo(bar)baz(blim)", the output should be
# reverseInParentheses(inputString) = "foorabbazmilb";

# For inputString = "foo(bar(baz))blim", the output should be
# reverseInParentheses(inputString) = "foobazrabblim".
# Because "foo(bar(baz))blim" becomes "foo(barzab)blim" and then "foobazrabblim".

# inputString = "foo(bar)baz", the output should be 
# reverseInParentheses(inputString) = "foorabbaz";

#### Using the python "eval" function..
#### Related video: https://realpython.com/lessons/python-eval-overview/
#### eval is a python built-in function that evaluates expressions


####

#s = "foo(bar)baz"
#s = "ab(cd)e"
#print(s[4:]) # start at position X and print the rest
#print(s[:4]) # start at beginning and print up and through position X
#print(s[2:4]) # start to right of X and print up and through Y

#Great site: https://pythex.org/

# regex notes:
# /t[aeiou]  -> brackets mean a range of values
# [^a-zA-Z]    not!

# \([a-zA-Z]*\) 
# foo(bar)baz(blim)
# matches (bar) and (blim)

# Meta-characters , like white space
# special meta-characters that represent a type of character like..
# Decimal digits, whitespace, words, etc.
# [aeiou].[aeiou]    "." means any character, except newline

# Regex Anchors:
# 

import re #You can use this, just put it above the function def.

a_string = "this is a [ki_te] message"
result = re.search(r"\[([A-Za-z0-9_]+)\]", a_string)
print(result.group(1))
