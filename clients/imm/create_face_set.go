package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFaceSetRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

func (req *CreateFaceSetRequest) Invoke(client *sdk.Client) (resp *CreateFaceSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreateFaceSet", "imm", "")
	resp = &CreateFaceSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFaceSetResponse struct {
	responses.BaseResponse
	RequestId  common.String
	SetId      common.String
	Status     common.String
	Photos     common.Long
	CreateTime common.String
	ModifyTime common.String
}
