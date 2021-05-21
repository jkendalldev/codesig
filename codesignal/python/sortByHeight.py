#!/usr/local/bin/python3

# Some people are standing in a row in a park. 
# There are trees between them which cannot be moved. 
# Your task is to rearrange the people by their heights in a non-descending order 
# without moving the trees. People can be very tall!

# Example

# For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
# sortByHeight(a) = [-1, 150, 160, 170, -1, -1, 180, 190].

line = [-1, 150, 190, 170, -1, -1, 160, 180]

def sortByHeight(line):
    people = sorted([num for num in line if num != -1 ])  # list comphrension
    person = 0

    for count, height in enumerate(line):  
        if height != -1:
            line[count] = people[person]
            person += 1
    
    return(line)

print(sortByHeight(line))



             

