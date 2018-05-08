package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AttachInstanceRamRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AttachInstanceRamRoleRequest) Invoke(client *sdk.Client) (resp *AttachInstanceRamRoleResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachInstanceRamRole", "ecs", "")
	resp = &AttachInstanceRamRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachInstanceRamRoleResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	TotalCount                   common.Integer
	FailCount                    common.Integer
	RamRoleName                  common.String
	AttachInstanceRamRoleResults AttachInstanceRamRoleAttachInstanceRamRoleResultList
}

type AttachInstanceRamRoleAttachInstanceRamRoleResult struct {
	InstanceId common.String
	Success    bool
	Code       common.String
	Message    common.String
}

type AttachInstanceRamRoleAttachInstanceRamRoleResultList []AttachInstanceRamRoleAttachInstanceRamRoleResult

func (list *AttachInstanceRamRoleAttachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachInstanceRamRoleAttachInstanceRamRoleResult)
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
