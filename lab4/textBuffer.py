import os

class TextBuffer:
  def __init__(self):
    self.text = ""
  
  def readFromFile(self, path: str):
    if not os.path.isfile(path):
      raise Exception(f"File with path {path} does not exist!")

    file = open(path, "r")
    self.text = file.read()
    file.close()
    