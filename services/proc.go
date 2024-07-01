package services

import (
	"github.com/shirou/gopsutil/v4/process"
)

func NewProcService() *ProcService {
	return &ProcService{}
}

type ProcService struct {
}

type ProcStat struct {
	PID            int32
	PPID           int32
	User           string
	VirtualMemory  uint64
	ResidentMemory uint64
	CPUPercent     float64
	MemPercent     float32
	Exe            string
}

func (ps *ProcService) Processes() ([]ProcStat, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	stats := make([]ProcStat, 0, len(procs))
	for _, proc := range procs {
		exe, _ := proc.Cmdline()
		ppid, _ := proc.Ppid()
		user, _ := proc.Username()
		mem, _ := proc.MemoryInfo()
		cpuPercent, _ := proc.CPUPercent()
		memPercent, _ := proc.MemoryPercent()
		stats = append(stats, ProcStat{
			PID:            proc.Pid,
			PPID:           ppid,
			User:           user,
			VirtualMemory:  mem.VMS,
			ResidentMemory: mem.RSS,
			CPUPercent:     cpuPercent,
			MemPercent:     memPercent,
			Exe:            exe,
		})
	}

	return stats, nil
}
