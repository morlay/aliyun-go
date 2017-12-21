package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMasterSlaveServerGroupsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeMasterSlaveServerGroupsRequest) Invoke(client *sdk.Client) (response *DescribeMasterSlaveServerGroupsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMasterSlaveServerGroupsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroups", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMasterSlaveServerGroupsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeMasterSlaveServerGroupsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMasterSlaveServerGroupsResponse struct {
	RequestId               string
	MasterSlaveServerGroups DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	MasterSlaveServerGroupId   string
	MasterSlaveServerGroupName string
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList []DescribeMasterSlaveServerGroupsMasterSlaveServerGroup

func (list *DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupsMasterSlaveServerGroup)
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
