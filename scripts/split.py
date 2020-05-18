import os

with open(os.path.dirname(os.path.realpath(__file__)) + "/run_all.sh", "r") as f:
    content = f.readlines()
    content = content[1:]
    content = content[:-1]

fileIndex = 0

while len(content) > 0:
    linesToWrite = content[:10]
    f = open(os.path.dirname(os.path.realpath(__file__)) +
             "/run_all_"+str(fileIndex)+".sh", "w+")
    f.write("#!/bin/bash\n")
    for line in linesToWrite:
        f.write(line)
    f.write("wait")
    content = content[10:]
    fileIndex += 1
    f.close()
    os.chmod(f.name, 0o755)

f = open(os.path.dirname(os.path.realpath(__file__)) + "/../Procfile", "a+")
for i in range(fileIndex):
    f.write("all_"+str(i)+": ""./scripts/run_all_"+str(i)+".sh\n")

with open(os.path.dirname(os.path.realpath(__file__)) + "/goz.sh", "r") as f:
    content = f.readlines()
    content = content[1:]
    content = content[:-1]

fileIndex = 0

while len(content) > 0:
    linesToWrite = content[:10]
    f = open(os.path.dirname(os.path.realpath(__file__)) +
             "/goz_"+str(fileIndex)+".sh", "w+")
    f.write("#!/bin/bash\n")
    for line in linesToWrite:
        f.write(line)
    f.write("wait")
    content = content[10:]
    fileIndex += 1
    f.close()
    os.chmod(f.name, 0o755)


f = open(os.path.dirname(os.path.realpath(__file__)) + "/../Procfile", "a+")
for i in range(fileIndex):
    f.write("goz_"+str(i)+": ""./scripts/goz_"+str(i)+".sh\n")
