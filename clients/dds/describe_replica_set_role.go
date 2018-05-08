package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicaSetRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeReplicaSetRoleRequest) Invoke(client *sdk.Client) (resp *DescribeReplicaSetRoleResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeReplicaSetRole", "dds", "")
	resp = &DescribeReplicaSetRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicaSetRoleResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DBInstanceId common.String
	ReplicaSets  DescribeReplicaSetRoleReplicaSetList
}

type DescribeReplicaSetRoleReplicaSet struct {
	ReplicaSetRole   common.String
	ConnectionDomain common.String
	ConnectionPort   common.String
	ExpiredTime      common.String
	NetworkType      common.String
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
