from imageViewer import ImageViewer
from viewerCreator import ViewerCreator

# exercise 1 
viewerCreator = ViewerCreator()

textViewer = viewerCreator.createViewer(r"./essay.txt")
textViewer.view()
stats = textViewer.getData()
print(f"numberOfLines: {stats.numberOfLines}")
print(f"numberOfWords: {stats.numberOfWords}")
print(f"numberOfNonalpha: {stats.numberOfNonalpha}")

imgViewer = viewerCreator.createViewer(r"./shampoo.jpeg")
imgViewer.view()