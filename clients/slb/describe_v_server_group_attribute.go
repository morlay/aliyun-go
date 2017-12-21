package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVServerGroupAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeVServerGroupAttributeRequest) Invoke(client *sdk.Client) (response *DescribeVServerGroupAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVServerGroupAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeVServerGroupAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVServerGroupAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVServerGroupAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVServerGroupAttributeResponse struct {
	RequestId        string
	VServerGroupId   string
	VServerGroupName string
	BackendServers   DescribeVServerGroupAttributeBackendServerList
}

type DescribeVServerGroupAttributeBackendServer struct {
	ServerId string
	Port     int
	Weight   int
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
