from imageViewer import ImageViewer
from viewerCreator import ViewerCreator
from multipleAccumulate import MultipleAccumulate

### exercise 1 ###
viewerCreator = ViewerCreator()

textViewer = viewerCreator.createViewer(r"./essay.txt")
# textViewer.view()
# stats = textViewer.getData()
# print(f"numberOfLines: {stats.numberOfLines}")
# print(f"numberOfWords: {stats.numberOfWords}")
# print(f"numberOfNonalpha: {stats.numberOfNonalpha}")

imgViewer = viewerCreator.createViewer(r"./shampoo.jpeg")
# imgViewer.view()

### exercise 2 ###
def multiply(a, b):
  return a * b

multipleAccumulate = MultipleAccumulate([1, 2, 3, 4, 5], lambda a, b: a + b, multiply, lambda a, b: max(a, b))
dictionary = multipleAccumulate.getData()
# print(dictionary)

### duck typing ###

for data in [multipleAccumulate, textViewer]:
  print(data.getData())