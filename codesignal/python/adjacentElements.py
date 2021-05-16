#!/usr/local/bin/python3

def adjacentElementsProduct(inputList):
  productList = []
  for i in range(len(inputList)):
    if i != (len(inputList) -1):
      n = i + 1
      p = (inputList[i] * inputList[n])
      productList.append(p)
  return(max(productList))

inputList = [1, 2, 3, 4, 5, 6]
print(adjacentElementsProduct(inputList))


 
 



