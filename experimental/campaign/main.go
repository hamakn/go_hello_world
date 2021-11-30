package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hamakn/go_hello_world/experimental/campaign/entity"
	v1 "github.com/hamakn/go_hello_world/experimental/campaign/v1"
)

var (
	campaignCandidates = []*entity.CampaignCandidate{
		v1.Candidate,
	}
)

func main() {
	now := time.Now()
	ctx := context.Background()
	cond := &entity.CampaignCondition{now}
	c := selectCampaign(ctx, cond)
	fmt.Println(c)
}

func selectCampaign(ctx context.Context, cond *entity.CampaignCondition) *entity.Campaign {
	for _, candidate := range campaignCandidates {
		if candidate.IsTarget(ctx, cond) {
			return candidate.Get(ctx, cond)
		}
	}
	return nil
}
