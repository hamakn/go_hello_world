package v1

import (
	"context"
	"time"

	"github.com/hamakn/go_hello_world/experimental/campaign/entity"
	"github.com/hamakn/go_hello_world/experimental/campaign/helper"
)

var (
	startedAt = time.Date(2019, 10, 2, 0, 0, 0, 0, time.UTC)
	endedAt   = time.Date(2019, 10, 5, 0, 0, 0, 0, time.UTC)

	myCampaign = &entity.Campaign{"This is my campaign"}

	Candidate = &entity.CampaignCandidate{
		IsTarget: func(ctx context.Context, cond *entity.CampaignCondition) bool {
			return helper.IsBetweenFunc(startedAt, endedAt)(cond.Now)
		},
		Get: func(ctx context.Context, cond *entity.CampaignCondition) *entity.Campaign {
			return myCampaign
		},
	}
)
