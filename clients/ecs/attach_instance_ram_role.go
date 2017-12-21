package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachInstanceRamRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AttachInstanceRamRoleRequest) Invoke(client *sdk.Client) (response *AttachInstanceRamRoleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachInstanceRamRoleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachInstanceRamRole", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AttachInstanceRamRoleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachInstanceRamRoleResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachInstanceRamRoleResponse struct {
	RequestId                    string
	TotalCount                   int
	FailCount                    int
	RamRoleName                  string
	AttachInstanceRamRoleResults AttachInstanceRamRoleAttachInstanceRamRoleResultList
}

type AttachInstanceRamRoleAttachInstanceRamRoleResult struct {
	InstanceId string
	Success    bool
	Code       string
	Message    string
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
