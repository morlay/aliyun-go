package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetFaceSetRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *GetFaceSetRequest) Invoke(client *sdk.Client) (resp *GetFaceSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetFaceSet", "imm", "")
	resp = &GetFaceSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFaceSetResponse struct {
	responses.BaseResponse
	RequestId  common.String
	SetId      common.String
	Status     common.String
	Photos     common.Long
	CreateTime common.String
	ModifyTime common.String
}
