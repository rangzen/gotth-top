package services

import "github.com/shirou/gopsutil/v4/host"

func NewGeneralService() *GeneralService {
	return &GeneralService{}
}

type GeneralService struct {
}

type GeneralStat struct {
	Hostname string
}

func (gs *GeneralService) GeneralStat() (GeneralStat, error) {
	info, err := host.Info()
	if err != nil {
		return GeneralStat{}, err
	}

	stats := GeneralStat{
		Hostname: info.Hostname,
	}
	return stats, nil
}
