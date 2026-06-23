package service

import (
	"loggin/internal/model"
	"loggin/internal/repository"
	"github.com/nxadm/tail"
)

type LogService struct {
	repo *repository.LogRepository
}

func NewFileService(r *repository.LogRepository) *LogService{
	return &LogService{
		repo: r,
	}
}

func logContent(log *model.Log, logs *tail.Tail, num *int){
	for line := range logs.Lines{
		if line.Num > *num || num == nil {
			item := model.LogItem{
				Order		: line.Num,
				TimeStamp	: line.Time,
				Line		: line.Text,
			}
			
			log.Content = append(log.Content, item)
		}
	}
}

func (l *LogService) GetLogContent(path string, num int) (*model.Log, error) {
	logs, err := repository.NewLogRepository().GetLog(path)
	
	if err != nil {
		return nil, err
	}

	var log model.Log

	logContent(&log, logs, &num)

	return &log, nil
}