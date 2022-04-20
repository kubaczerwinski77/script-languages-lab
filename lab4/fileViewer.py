from abc import ABC, abstractmethod

class FileViewer(ABC):
  def __init__(self, path: str):
    super().__init__()
    self.path = path

  @abstractmethod
  def view(self):
    pass
