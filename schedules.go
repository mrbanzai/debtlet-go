package main

import (
	"database/sql"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/teambition/rrule-go"
	"golang.org/x/exp/slices"
)

var (
	SemiMonthly, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.MONTHLY,
		Bysetpos:   []int{5, -1},
		Bymonthday: []int{9, 10, 11, 12, 13, 14, 15, 26, 27, 28, 29, 30, 31},
		Byweekday:  []rrule.Weekday{rrule.MO, rrule.TU, rrule.WE, rrule.TH, rrule.FR},
	})
	BiWeekly, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.WEEKLY,
		Interval: 2,
	})
)

type Event struct {
	Transaction Transaction
	Recurring   bool
	Recurrence  *rrule.RRule
	End         sql.NullTime
	Overrides   EventOverrides
}

type EventOverride struct {
	DateTime time.Time
	Amount   *money.Money
	Delete   bool
}

type EventOverrides []EventOverride

func (o EventOverrides) Sort() {
	slices.SortFunc(o, func(a EventOverride, b EventOverride) bool {
		return a.DateTime.Before(b.DateTime)
	})
}
