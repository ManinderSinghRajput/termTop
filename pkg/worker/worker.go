package worker

import (
	"volthatop/pkg/cpu"
	"volthatop/pkg/disk"
	"volthatop/pkg/mem"
	"volthatop/pkg/net"
	"volthatop/pkg/temp"
	"sync"
)

type Status struct {
	CPULoad *cpu.CPULoad
	CPUFreq *cpu.CPUFreq
	Mem     *mem.MemStat
	Temp    *temp.Temp
	Disk    *[]disk.DiskInfo
	Net     *net.NetStat
}

func Worker() *Status {

	var status Status

	var wg sync.WaitGroup
	wg.Add(6)

	//CPU Load
	go func() {
		defer wg.Done()
		status.CPULoad = cpu.ComputeCPULoad()
	}()

	// CPU Frequency
	go func() {
		defer wg.Done()
		status.CPUFreq = cpu.ExtractCPUFrequency()
	}()

	//Mem Stats
	go func() {
		defer wg.Done()
		status.Mem = mem.GetMemStats()
	}()

	// Temp
	go func() {
		defer wg.Done()
		status.Temp = temp.ExtractTemp()
	}()

	// Disk
	go func() {
		defer wg.Done()
		status.Disk = disk.ExtractDiskUsage()
	}()

	// Net
	go func() {
		defer wg.Done()
		status.Net = net.ComputeNetStats()
	}()
	wg.Wait()

	return &status
}
