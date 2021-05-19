#!/usr/local/bin/python3

# this problem makes no sense.

#return sum(min(s1.count(x), s2.count(x)) for x in set(s1))

# For s1 = "a a b c c" and s2 = "a d c a a", the output should be
#          "a d c a a"


#s1: "a b c d e f g h x y z t t w"
#s2: "h g f e d c b a a b c w w t"

# a, w, t, e, h, d, 
# 10


#commonCharacterCount(s1, s2) = 3.

s1 = "a"
s2 = "a"

com = [min(s1.count(i),s2.count(i)) for i in set(s1)]

# {'a', 'c', 'b'}
#print(min(s1.count("a"),s2.count("a"))) # 2
#print(min(s1.count("c"),s2.count("c"))) # 1
#print(min(s1.count("b"),s2.count("b"))) # 0

 
print(sum(com))

#print(min(5,2,1,80)) 


#Strings have 3 common characters - 2 "a"s and 1 "c".
#def commonCharacterCount(s1, s2):

#def commonCharacterCount(s1, s2):
    #com = [min(s1.count(i),s2.count(i)) for i in set(s1)]
    #return sum(com)