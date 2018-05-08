package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCenAttachedChildInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCenAttachedChildInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeCenAttachedChildInstancesResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenAttachedChildInstances", "cbn", "")
	resp = &DescribeCenAttachedChildInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenAttachedChildInstancesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalCount     common.Integer
	PageNumber     common.Integer
	PageSize       common.Integer
	ChildInstances DescribeCenAttachedChildInstancesChildInstanceList
}

type DescribeCenAttachedChildInstancesChildInstance struct {
	CenId                 common.String
	ChildInstanceId       common.String
	ChildInstanceType     common.String
	ChildInstanceRegionId common.String
	ChildInstanceOwnerId  common.Long
	Status                common.String
}

type DescribeCenAttachedChildInstancesChildInstanceList []DescribeCenAttachedChildInstancesChildInstance

func (list *DescribeCenAttachedChildInstancesChildInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenAttachedChildInstancesChildInstance)
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
