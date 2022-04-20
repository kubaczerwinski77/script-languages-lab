from textViewer import TextViewer
from imageViewer import ImageViewer
import mimetypes

class ViewerCreator:
  def __init__(self):
    pass

  def detectViewerType(self, path: str):
    # guess_type return a tuple (type, encoding)
    fileType = mimetypes.guess_type(path)[0].split("/")[0]

    if fileType == "image":
      return ImageViewer
    elif fileType == "text":
      return TextViewer
    else:
      raise Exception(f"Invalid file type. Image or text required! Given type: {fileType}")


  def createViewer(self, path: str):
    return self.detectViewerType(path)(path)