package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingActivitiesRequest struct {
	requests.RpcRequest
	ScalingActivityId9   string `position:"Query" name:"ScalingActivityId.9"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingActivityId5   string `position:"Query" name:"ScalingActivityId.5"`
	ScalingActivityId6   string `position:"Query" name:"ScalingActivityId.6"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ScalingActivityId7   string `position:"Query" name:"ScalingActivityId.7"`
	ScalingActivityId8   string `position:"Query" name:"ScalingActivityId.8"`
	ScalingActivityId1   string `position:"Query" name:"ScalingActivityId.1"`
	ScalingActivityId2   string `position:"Query" name:"ScalingActivityId.2"`
	ScalingActivityId3   string `position:"Query" name:"ScalingActivityId.3"`
	ScalingActivityId4   string `position:"Query" name:"ScalingActivityId.4"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	StatusCode           string `position:"Query" name:"StatusCode"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ScalingActivityId11  string `position:"Query" name:"ScalingActivityId.11"`
	ScalingActivityId10  string `position:"Query" name:"ScalingActivityId.10"`
	ScalingActivityId13  string `position:"Query" name:"ScalingActivityId.13"`
	ScalingActivityId12  string `position:"Query" name:"ScalingActivityId.12"`
	ScalingActivityId15  string `position:"Query" name:"ScalingActivityId.15"`
	ScalingActivityId14  string `position:"Query" name:"ScalingActivityId.14"`
	ScalingActivityId17  string `position:"Query" name:"ScalingActivityId.17"`
	ScalingActivityId16  string `position:"Query" name:"ScalingActivityId.16"`
	ScalingActivityId19  string `position:"Query" name:"ScalingActivityId.19"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingActivityId18  string `position:"Query" name:"ScalingActivityId.18"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingActivityId20  string `position:"Query" name:"ScalingActivityId.20"`
}

func (req *DescribeScalingActivitiesRequest) Invoke(client *sdk.Client) (resp *DescribeScalingActivitiesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingActivities", "ess", "")
	resp = &DescribeScalingActivitiesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingActivitiesResponse struct {
	responses.BaseResponse
	TotalCount        int
	PageNumber        int
	PageSize          int
	RequestId         string
	ScalingActivities DescribeScalingActivitiesScalingActivityList
}

type DescribeScalingActivitiesScalingActivity struct {
	ScalingActivityId   string
	ScalingGroupId      string
	Description         string
	Cause               string
	StartTime           string
	EndTime             string
	Progress            int
	StatusCode          string
	StatusMessage       string
	TotalCapacity       string
	AttachedCapacity    string
	AutoCreatedCapacity string
}

type DescribeScalingActivitiesScalingActivityList []DescribeScalingActivitiesScalingActivity

func (list *DescribeScalingActivitiesScalingActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingActivitiesScalingActivity)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
