package tasks

import (
	"atlas-fec/expression"
	"atlas-fec/kafka/producers"
	"context"
	"log"
	"time"
)

type ExpressionRevert struct {
	l        *log.Logger
	interval time.Duration
}

func NewExpressionRevert(l *log.Logger, interval time.Duration) *ExpressionRevert {
	l.Printf("[INFO] initializing expression revert task to run every %dms", interval.Milliseconds())
	return &ExpressionRevert{l, interval}
}

func (e *ExpressionRevert) Run() {
	for _, exp := range expression.GetCache().PopExpired() {
		producers.CharacterExpressionChanged(e.l, context.Background()).Emit(exp.CharacterId(), exp.MapId(), exp.Expression())
	}
}

func (e *ExpressionRevert) SleepTime() time.Duration {
	return e.interval
}
