package cron

import (
	"time"

	"github.com/robfig/cron/v3"
)

type SaveMaps struct{}

func (s SaveMaps) Run() {

}

func InitCron() *cron.Cron {
	// 시간대 설정 (KST)
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil
	}

	// cron 인스턴스 생성 (초 단위 지원 활성화)
	c := cron.New(cron.WithSeconds(), cron.WithLocation(loc))

	// 1분마다 실행
	_, err = c.AddJob("0 */1 * * * *", SaveMaps{})

	if err != nil {
		return nil
	}

	c.Start()

	return c
}
