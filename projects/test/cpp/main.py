import time

start_time = time.time()
x = 0
while x != 1000000000:
    v = x * 100
    v2 = v / 99
    x = x + 1
end_time = time.time()
print(end_time - start_time)