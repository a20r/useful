package assert

import "github.com/sirupsen/logrus"

type Op struct {
	log *logrus.Entry
}

func NewOp(l *logrus.Entry) Op {
	return Op{log: l}
}

func (o Op) NoError(err error) {
	if err != nil {
		o.log.WithError(err).Fatalf("🚨 Encountered unexpected error: %s", err.Error())
	}
}
