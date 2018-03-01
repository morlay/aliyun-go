package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	SetId      string
	Status     string
	Photos     int64
	CreateTime string
	ModifyTime string
}
