package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateTagSetRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

func (req *CreateTagSetRequest) Invoke(client *sdk.Client) (resp *CreateTagSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreateTagSet", "imm", "")
	resp = &CreateTagSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTagSetResponse struct {
	responses.BaseResponse
	RequestId  common.String
	SetId      common.String
	Status     common.String
	Photos     common.Long
	CreateTime common.String
	ModifyTime common.String
}
