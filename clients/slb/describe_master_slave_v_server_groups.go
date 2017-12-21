package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMasterSlaveVServerGroupsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeMasterSlaveVServerGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeMasterSlaveVServerGroupsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveVServerGroups", "slb", "")
	resp = &DescribeMasterSlaveVServerGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMasterSlaveVServerGroupsResponse struct {
	responses.BaseResponse
	RequestId                string
	MasterSlaveVServerGroups DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup struct {
	MasterSlaveVServerGroupId   string
	MasterSlaveVServerGroupName string
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList []DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup

func (list *DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup)
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
