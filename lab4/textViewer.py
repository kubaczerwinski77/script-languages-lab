from textBuffer import TextBuffer
from fileViewer import FileViewer
from textStats import TextStats
import subprocess

class TextViewer(TextBuffer, FileViewer):
  def __init__(self, path: str):
    FileViewer.__init__(self, path)
    TextBuffer.__init__(self)

    self.readFromFile(path)

    # previous was:
    # self.stats = TextStats().compute(self.text)
    # but it did not work
    self.stats = TextStats()
    self.stats.compute(self.text)

  def getData(self):
    return self.stats

  def view(self):
    subprocess.Popen(["open", self.path])
  