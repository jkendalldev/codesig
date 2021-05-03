#!/usr/local/bin/python3

#def almostIncreasingSequence(sequence):
    #f1 = sum([1 for a, b in zip(sequence[:-1], sequence[1:]) if a>=b ]) <= 1 
    #f2 = sum([1 for a, c in zip(sequence[:-2], sequence[2:]) if a>=c ]) <= 1
    #return f1 and f2

# sequence = [1, 2, 3, 4]

############################## GOOD STUFF
# x = sum([1 for a, b in [(1, 9), (1, 8)] if a>=b ])  <= 1  # I get what this does now!
                                                          # It simply sets x to True
                                                          # IF A is > B only 1 or less times for each
                                                          # number set iterated through the loop!
                                                          # It's just a fancy counter really.

############## SO, THE ONLY NEXT THING TO DO IS, UNDERSTAND THE LOGIC BEHIND WHY WE ARE DOING THE ZIP
# STUFF THE WAY WE ARE DOING IT...
# zip(sequence[:-1], sequence[1:])





