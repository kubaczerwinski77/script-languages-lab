from functools import reduce

class MultipleAccumulate:
  def __init__(self, dataList: list, *functions):
    self.dataList = dataList
    for elem in functions:
      if not callable(elem):
        raise Exception("All args must be callable (apart from dataList)!")
    self.functions = functions
  
  def getData(self):
    lambdaCouter = 1
    dictionary = {}

    for func in self.functions:
      funcName = func.__name__
      if funcName == "<lambda>":
        funcName = f"lambda{lambdaCouter}"
        lambdaCouter += 1
      dictionary[funcName] = reduce(func, self.dataList)
    
    return dictionary
