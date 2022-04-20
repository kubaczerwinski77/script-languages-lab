class TextStats:
  def __init__(self):
    self.numberOfLines = 0
    self.numberOfWords = 0
    self.numberOfNonalpha = 0

  def __str__(self):
    return f"numberOfLines: {self.numberOfLines}\nnumberOfWords: {self.numberOfWords}\nnumberOfNonalpha: {self.numberOfNonalpha}"

  def compute(self, text: str):
    self.numberOfLines = len(text.splitlines())
    self.numberOfWords = len(text.split())
    self.numberOfNonalpha = (lambda text: sum(1 - char.isalpha() for char in text))(text)