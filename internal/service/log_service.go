package service

import (
	"loggin/internal/model"
	"loggin/internal/repository"
)

type LogService struct {
	repo *repository.LogRepository
}

func NewFileService(r *repository.LogRepository) *LogService{
	return &LogService{
		repo: r,
	}
}

func (l *LogService) GetLogContent(path string) (*model.Log, error) {
	logs, err := repository.NewLogRepository().GetLog(path)
	
	if err != nil {
		return nil, err
	}

	var log model.Log

	for line := range logs.Lines{
		item := model.LogItem{
			Order		: line.Num,
			TimeStamp	: line.Time,
			Line		: line.Text,
		}
		
		log.Content = append(log.Content, item)
	}

	return &log, nil
}

func (l *LogService) GetLogContentByNum(path string, num int) (*model.Log, error) {
	logs, err := repository.NewLogRepository().GetLog(path)
	
	if err != nil {
		return nil, err
	}

	var log model.Log

	for line := range logs.Lines{
		if line.Num > num {
			item := model.LogItem{
				Order		: line.Num,
				TimeStamp	: line.Time,
				Line		: line.Text,
			}
			
			log.Content = append(log.Content, item)
		}
	}

	return &log, nil
}