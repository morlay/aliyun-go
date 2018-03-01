package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetTagSetRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *GetTagSetRequest) Invoke(client *sdk.Client) (resp *GetTagSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetTagSet", "imm", "")
	resp = &GetTagSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetTagSetResponse struct {
	responses.BaseResponse
	RequestId  string
	SetId      string
	Status     string
	CreateTime string
	ModifyTime string
	Photos     int64
}
