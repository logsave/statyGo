import ctypes

lib = ctypes.CDLL('./sum.so')
print(lib.sum(1,5))
