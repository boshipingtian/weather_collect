package crontab

type CWork interface {
	getName() string
	Run()
}

var WorkMap = map[string]CWork{}

func PutWork(work CWork) {
	if _, ok := WorkMap[work.getName()]; !ok {
		WorkMap[work.getName()] = work
	}
}
