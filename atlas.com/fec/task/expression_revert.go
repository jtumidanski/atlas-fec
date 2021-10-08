package tasks

import (
	"atlas-fec/expression"
	"atlas-fec/kafka/producers"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"time"
)

const ExpressionRevertTask = "expression_revert_task"

type ExpressionRevert struct {
	l        logrus.FieldLogger
	interval time.Duration
}

func NewExpressionRevert(l logrus.FieldLogger, interval time.Duration) *ExpressionRevert {
	l.Infof("Initializing expression revert task to run every %dms", interval.Milliseconds())
	return &ExpressionRevert{l, interval}
}

func (e *ExpressionRevert) Run() {
	span := opentracing.StartSpan(ExpressionRevertTask)
	for _, exp := range expression.GetCache().PopExpired() {
		producers.CharacterExpressionChanged(e.l, span)(exp.CharacterId(), exp.MapId(), exp.Expression())
	}
	span.Finish()
}

func (e *ExpressionRevert) SleepTime() time.Duration {
	return e.interval
}
