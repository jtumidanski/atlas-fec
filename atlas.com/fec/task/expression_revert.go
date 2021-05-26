package tasks

import (
	"atlas-fec/expression"
	"atlas-fec/kafka/producers"
	"github.com/sirupsen/logrus"
	"time"
)

type ExpressionRevert struct {
	l        logrus.FieldLogger
	interval time.Duration
}

func NewExpressionRevert(l logrus.FieldLogger, interval time.Duration) *ExpressionRevert {
	l.Infof("Initializing expression revert task to run every %dms", interval.Milliseconds())
	return &ExpressionRevert{l, interval}
}

func (e *ExpressionRevert) Run() {
	for _, exp := range expression.GetCache().PopExpired() {
		producers.CharacterExpressionChanged(e.l)(exp.CharacterId(), exp.MapId(), exp.Expression())
	}
}

func (e *ExpressionRevert) SleepTime() time.Duration {
	return e.interval
}
