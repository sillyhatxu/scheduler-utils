package scheduler

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Message1 struct {
	Name string
}

func (m Message1) Execute() {
	logrus.Infof("name : %s", m.Name)
}

type Message2 struct {
	Name string
}

func (m Message2) Execute() {
	logrus.Infof("name : %s start %v", m.Name, time.Now().Format("2006-01-02T15:04:05"))
	time.Sleep(5 * time.Second)
	logrus.Infof("name : %s end %v", m.Name, time.Now().Format("2006-01-02T15:04:05"))
}

type Message3 struct {
	Name string
}

func (m Message3) Execute() {
	logrus.Infof("name : %s start %v", m.Name, time.Now().Format("2006-01-02T15:04:05"))
	time.Sleep(5 * time.Second)
	logrus.Infof("name : %s end %v", m.Name, time.Now().Format("2006-01-02T15:04:05"))
}

func TestMessage1(t *testing.T) {
	schedulerMessage1, err := NewScheduler(&Message3{Name: "name-1"})
	assert.Nil(t, err)
	schedulerMessage1.Run()
	select {}
}
func TestMessage2(t *testing.T) {
	schedulerMessage2, err := NewScheduler(&Message3{Name: "name-2"}, Start("00:00:00"), Interval(2*time.Second), Concurrent(false))
	assert.Nil(t, err)
	schedulerMessage2.Run()
	select {}
}

func TestMessage3(t *testing.T) {
	schedulerMessage3, err := NewScheduler(&Message3{Name: "name-3"}, Start("00:00:00"), Interval(2*time.Second))
	assert.Nil(t, err)
	schedulerMessage3.Run()
	select {}
}
