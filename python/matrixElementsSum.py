#!/usr/local/bin/python3

def matrixElementsSum(floors):
    cost_of_rooms = 0
    cols = [list(x) for x in zip(*floors)]
    for count, col in enumerate(cols):  
        for cnt, room in enumerate(col):
            if cnt != 0:
                rooms_above_me = col[:cnt]
                if 0 not in rooms_above_me:
                    cost_of_rooms += room
    rooms_total = cost_of_rooms + sum(floors[0])
    return rooms_total
floors = [[10, 0, 10, 0], [1, 5, 0, 1], [2, 1, 3, 10]]
print(matrixElementsSum(floors))

