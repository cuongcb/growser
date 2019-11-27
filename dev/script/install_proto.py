import subprocess
import sys

protoURL = "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.0/protoc-3.11.0-linux-x86_64.zip"
protoFile = "proto.zip"
srcDir = sys.argv[1]
absFilePath = srcDir + "/" + protoFile
outDir = sys.argv[2]

def installProto():
    subprocess.run(["rm", "-vrf", outDir])
    subprocess.run(["wget", protoURL, "-O", absFilePath])
    subprocess.run(["unzip", absFilePath, "-d", outDir])

def cleanProto():
    subprocess.run(["rm", "-vrf", outDir])
    subprocess.run(["rm", "-vrf", absFilePath])

def main():
    if (len(sys.argv) > 3 and sys.argv[3] == "clean"):
        cleanProto()
    else:
        installProto()

if __name__ == '__main__':
    main()
