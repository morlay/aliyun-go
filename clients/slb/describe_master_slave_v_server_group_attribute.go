package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMasterSlaveVServerGroupAttributeRequest struct {
	requests.RpcRequest
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveVServerGroupId string `position:"Query" name:"MasterSlaveVServerGroupId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
}

func (req *DescribeMasterSlaveVServerGroupAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeMasterSlaveVServerGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveVServerGroupAttribute", "slb", "")
	resp = &DescribeMasterSlaveVServerGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMasterSlaveVServerGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId                   string
	MasterSlaveVServerGroupId   string
	MasterSlaveVServerGroupName string
	MasterSlaveBackendServers   DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId string
	Port     int
	Weight   int
	IsBackup int
	Type     string
	ServerIp string
	VpcId    string
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer)
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
