package expression

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"time"
)

const RevertTaskName = "expression_revert_task"

type RevertTask struct {
	l        logrus.FieldLogger
	interval time.Duration
}

func NewRevertTask(l logrus.FieldLogger, interval time.Duration) *RevertTask {
	l.Infof("Initializing expression revert task to run every %dms", interval.Milliseconds())
	return &RevertTask{l, interval}
}

func (e *RevertTask) Run() {
	span := opentracing.StartSpan(RevertTaskName)
	for _, exp := range getCache().popExpired() {
		emitExpressionChanged(e.l, span)(exp.CharacterId(), exp.MapId(), exp.Expression())
	}
	span.Finish()
}

func (e *RevertTask) SleepTime() time.Duration {
	return e.interval
}
