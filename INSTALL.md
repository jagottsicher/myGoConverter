# Installation

## Binaries/Releases
Download the binary from the [releases](https://github.com/jagottsicher/myGoConverter/releases) page matching your operation system. install in the correstponding path:
Reccomendation to use as single user 
### Linux/U**X (/bin/bash)

### macOS (Terminal)

### Windows (cmd.exe/Powershell)

### Check version
Execute ```turn --version``` on one of your operation system's shells.

## Build from Source

### Prerequisities

### Linux/U**X (/bin/bash)

### macOS (Terminal)

### Windows (cmd.exe/Powershell)

from here stick to the instruction as described in the section Binaries/Releases (see above)



$ env GOOS=linux GOARCH=amd64 go build -o ../bin/linux/deployer  

$ env GOOS=windows GOARCH=amd64 go build -o ../bin/windows/deployer.exe  

$ env GOOS=darwin GOARCH=amd64 go build -o ../bin/darwin/deployer  

