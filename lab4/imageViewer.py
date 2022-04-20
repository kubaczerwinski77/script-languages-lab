from fileViewer import FileViewer
import subprocess

class ImageViewer(FileViewer):
  def __init__(self, path: str):
    super().__init__(path)

  def view(self):
    subprocess.Popen(["open", self.path])
