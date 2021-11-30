package entity

import (
	"context"
	"time"
)

type Campaign struct {
	Name string
}

type CampaignCondition struct {
	Now time.Time
}

type CampaignCandidate struct {
	IsTarget func(ctx context.Context, cond *CampaignCondition) bool
	Get      func(ctx context.Context, cond *CampaignCondition) *Campaign
}
