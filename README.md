# simpledocker

### 目录挂载
```shell
/root/imageName
/root/writeLayer/containerName
/root/mnt/containerName
/var/run/go-docker/containerName/config.json
/var/run/go-docker/containerName/container.log


/var/lib/simpledocker
    /containerName
        /lower
        /upper
        /work
        /merged

/var/run/simpledocker/containerName/config.json
/var/run/simpledocker/containerName/container.log
```

```go
type Info struct {
    Id   string
    Name string
	Env	string
	Cmd	[]string
    State struct {
        Status  string    // 容器状态
        Pid     int    // 容器Pid
        StartedAt string  // 最新的启动时间
        ExitCode  int  // 上一次停止状态
        FinishedAt string // 上一次停止日期
    }
    Mount []string		// 挂载目录 /path:path
    GraphDriver struct{
        Type		string
        ReadonlyDir string	// 容器只读层路径
        WriteDir	string	// 读写层
        WorkDir		string	// 挂载层
    }
    Network struct{}
}
```