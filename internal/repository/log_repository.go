package repository

import(
	"github.com/nxadm/tail"
)

type LogRepository struct {}

func NewLogRepository() *LogRepository{
    return &LogRepository{}
}

func (l *LogRepository) GetLog(path string) (*tail.Tail, error) {
    log, err := tail.TailFile(path, tail.Config{Follow: false})
    if err != nil {
        return nil, err
    }
    return log, nil
}