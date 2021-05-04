#!/usr/local/bin/python3

import random
import pdb #Python debugger

print('Guess a number between 1 and 10')
number = random.randint(1, 10)

pdb.set_trace()
guess = input()

if guess == number:
    print(f'Correct - the number was {number}')
else:
    print(f'Sorry - the number was {number}, not {guess}')

# Freeze time with pdb!
# once inside pdb!
# type "l" command to see where you are in the program once pdb takes over
# type "c" command to continue!
# type a variable name to get more info about it! rather than using print!
# so powerful!