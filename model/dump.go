package response

type DumpParam struct {
	Pid int32 `json:"pid" comment:"process ID"` //进程ID

	Dir string `json:"dir" comment:"Freeze the file saving address"` //冻结文件保存地址

	ShellJob bool `json:"shellJob" comment:"Is it a command line program"` //是否是一个命令行程序

	TcpEstablished bool `json:"tcpEstablished" comment:"Whether a TCP connection has been established"` //是否已建立TCP连接

	TcpSkipInFlight bool `json:"tcpSkipInFlight" comment:"Whether to skip a running TCP connection"` //是否跳过正在运行的TCP连接

	LeaveRunning bool `json:"leaveRunning" comment:"Freezes the current program process without stopping the program"` //在不停止程序时冻结当前程序进程

	TcpClose bool `json:"tcpClose" comment:"None example Close the established TCP"` // 关闭已建立的TCP
}
