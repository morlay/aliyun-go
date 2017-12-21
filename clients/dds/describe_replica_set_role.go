package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeReplicaSetRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeReplicaSetRoleRequest) Invoke(client *sdk.Client) (response *DescribeReplicaSetRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeReplicaSetRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeReplicaSetRole", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeReplicaSetRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeReplicaSetRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeReplicaSetRoleResponse struct {
	RequestId    string
	DBInstanceId string
	ReplicaSets  DescribeReplicaSetRoleReplicaSetList
}

type DescribeReplicaSetRoleReplicaSet struct {
	ReplicaSetRole   string
	ConnectionDomain string
	ConnectionPort   string
	ExpiredTime      string
	NetworkType      string
}

type DescribeReplicaSetRoleReplicaSetList []DescribeReplicaSetRoleReplicaSet

func (list *DescribeReplicaSetRoleReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaSetRoleReplicaSet)
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
