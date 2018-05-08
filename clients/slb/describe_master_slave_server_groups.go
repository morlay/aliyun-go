package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeMasterSlaveServerGroupsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeMasterSlaveServerGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeMasterSlaveServerGroupsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroups", "slb", "")
	resp = &DescribeMasterSlaveServerGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMasterSlaveServerGroupsResponse struct {
	responses.BaseResponse
	RequestId               common.String
	MasterSlaveServerGroups DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	MasterSlaveServerGroupId   common.String
	MasterSlaveServerGroupName common.String
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
