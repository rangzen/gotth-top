package services

import "github.com/shirou/gopsutil/v4/host"

func NewGeneralService() *GeneralService {
	return &GeneralService{}
}

type GeneralService struct {
}

type GeneralStat struct {
	Hostname             string
	Uptime               uint64
	BootTime             uint64
	Procs                uint64
	OS                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	KernelVersion        string
	KernelArch           string
	VirtualizationSystem string
	VirtualizationRole   string
	HostID               string
}

func (gs *GeneralService) GeneralStat() (GeneralStat, error) {
	info, err := host.Info()
	if err != nil {
		return GeneralStat{}, err
	}

	stats := GeneralStat{
		Hostname:             info.Hostname,
		Uptime:               info.Uptime,
		BootTime:             info.BootTime,
		Procs:                info.Procs,
		OS:                   info.OS,
		Platform:             info.Platform,
		PlatformFamily:       info.PlatformFamily,
		PlatformVersion:      info.PlatformVersion,
		KernelVersion:        info.KernelVersion,
		KernelArch:           info.KernelArch,
		VirtualizationSystem: info.VirtualizationSystem,
		VirtualizationRole:   info.VirtualizationRole,
		HostID:               info.HostID,
	}
	return stats, nil
}
