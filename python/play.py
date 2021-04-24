#!/usr/local/bin/python3
def digitsManipulations(n):
  number_string = str(n)
  number_map = map(int, number_string)
  number_list = list(number_map)
  product = 1
  for num in number_list:
    product = product * num
  #print(product)

  #sum = 1
  mysum = sum(number_list)
  #print(mysum)

  difference = product - mysum
  #print(difference)
  return(difference)

 

print(digitsManipulations(1234))

#my answer
def digitsManipulations(n):
  number_string = str(n)
  number_map = map(int, number_string)
  number_list = list(number_map)
  product = 1
  
  for num in number_list:
    product = product * num
    
  mysum = sum(number_list)
  difference = product - mysum
  
  return(difference)