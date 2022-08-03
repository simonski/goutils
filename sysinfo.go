package goutils

import (
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func (*SysInfo) IsLinux() bool {
	return strings.ToLower(runtime.GOOS) == "linux"
}

func (*SysInfo) IsWindows() bool {
	return strings.ToLower(runtime.GOOS) == "windows"
}

func (*SysInfo) IsMacOS() bool {
	return strings.ToLower(runtime.GOOS) == "darwin"
}

func NewSysInfo() *SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	// diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"
	diskStat, _ := disk.Usage("/") // If you're in Unix change this "\\" for "/"

	info := SysInfo{}

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	return &info

}
