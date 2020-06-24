import os
from random import randint
import time

viable = [num for num in range(1, 201)]

path = "Go"
subfolders = [ f.name.split(" ")[0] for f in os.scandir(path) if f.is_dir()]

for problem in subfolders:
    if problem.isnumeric():
        viable.remove(int(problem))

print("Today we will solve ", end="")
for _ in range(5):
    time.sleep(0.5)
    print(".", end="")
    
print("",viable[randint(0, len(viable))])
input("Press any key to exit.")