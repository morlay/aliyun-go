package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDeploymentSetsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	DeploymentSetIds     string `position:"Query" name:"DeploymentSetIds"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	PageSize             int    `position:"Query" name:"PageSize"`
	Strategy             string `position:"Query" name:"Strategy"`
}

func (req *DescribeDeploymentSetsRequest) Invoke(client *sdk.Client) (resp *DescribeDeploymentSetsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDeploymentSets", "ecs", "")
	resp = &DescribeDeploymentSetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDeploymentSetsResponse struct {
	responses.BaseResponse
	RequestId      string
	RegionId       string
	TotalCount     int
	PageNumber     int
	PageSize       int
	DeploymentSets DescribeDeploymentSetsDeploymentSetList
}

type DescribeDeploymentSetsDeploymentSet struct {
	DeploymentSetId          string
	DeploymentSetDescription string
	DeploymentSetName        string
	Strategy                 string
	Domain                   string
	Granularity              string
	InstanceAmount           int
	CreationTime             string
}

type DescribeDeploymentSetsDeploymentSetList []DescribeDeploymentSetsDeploymentSet

func (list *DescribeDeploymentSetsDeploymentSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetsDeploymentSet)
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
