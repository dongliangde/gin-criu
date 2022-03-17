package response

type RestoreParam struct {
	Dir string `json:"dir" comment:"freeze the file saving address"` //冻结文件保存地址

	ShellJob bool `json:"shellJob" comment:"is it a command line program"` //是否是一个命令行程序

	TcpEstablished bool `json:"tcpEstablished" comment:"whether a tcp connection has been established"` //是否已建立TCP连接

	TcpSkipInFlight bool `json:"tcpSkipInFlight" comment:"whether to skip a running tcp connection"` //是否跳过正在运行的TCP连接

	LeaveRunning bool `json:"leaveRunning" comment:"Freezes the current program process without stopping the program"` //是否跳过正在运行的TCP连接

	TcpClose bool `json:"tcpClose" comment:"none example close the established tcp"` // 关闭已建立的TCP

	RstSibling bool `json:"rstSibling" comment:"restore the root task to its sibling"` // 将根任务还原为同级
}
