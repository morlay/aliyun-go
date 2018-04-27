package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeGeographicRegionMembershipRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	GeographicRegionId   string `position:"Query" name:"GeographicRegionId"`
}

func (req *DescribeGeographicRegionMembershipRequest) Invoke(client *sdk.Client) (resp *DescribeGeographicRegionMembershipResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeGeographicRegionMembership", "cbn", "")
	resp = &DescribeGeographicRegionMembershipResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeGeographicRegionMembershipResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	RegionIds  DescribeGeographicRegionMembershipRegionIdList
}

type DescribeGeographicRegionMembershipRegionId struct {
	RegionId string
}

type DescribeGeographicRegionMembershipRegionIdList []DescribeGeographicRegionMembershipRegionId

func (list *DescribeGeographicRegionMembershipRegionIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGeographicRegionMembershipRegionId)
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
