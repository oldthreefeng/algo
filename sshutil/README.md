# ssh utils

## Usage

```go
package main

import (
    "github.com/oldthreefeng/algo/sshutil"
)

func main()  {
    s := sshutils.New()
    // Cmd is exec cmd on host
    s.Cmd(host, cmd)
    //CopyConfigFile is copy local file or bytes to remote file path
    s.CopyConfigFile(host, remoteFilePath , localFilePathOrBytes)
    // CopyRemoteFileToLocal is scp remote file to local
    s.CopyRemoteFileToLocal(host, localFilePath, remoteFilePath) 
}
```