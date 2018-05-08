package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	requests.RpcRequest
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
}

func (req *DescribeMasterSlaveServerGroupAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeMasterSlaveServerGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroupAttribute", "slb", "")
	resp = &DescribeMasterSlaveServerGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	MasterSlaveServerGroupId   common.String
	MasterSlaveServerGroupName common.String
	MasterSlaveBackendServers  DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId   common.String
	Port       common.Integer
	Weight     common.Integer
	ServerType common.String
	Type       common.String
	ServerIp   common.String
	VpcId      common.String
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer)
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
