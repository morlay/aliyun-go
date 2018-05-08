package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVServerGroupAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeVServerGroupAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeVServerGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeVServerGroupAttribute", "slb", "")
	resp = &DescribeVServerGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVServerGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	VServerGroupId   common.String
	VServerGroupName common.String
	BackendServers   DescribeVServerGroupAttributeBackendServerList
}

type DescribeVServerGroupAttributeBackendServer struct {
	ServerId common.String
	Port     common.Integer
	Weight   common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type DescribeVServerGroupAttributeBackendServerList []DescribeVServerGroupAttributeBackendServer

func (list *DescribeVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVServerGroupAttributeBackendServer)
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
