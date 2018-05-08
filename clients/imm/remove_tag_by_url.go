package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveTagByUrlRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *RemoveTagByUrlRequest) Invoke(client *sdk.Client) (resp *RemoveTagByUrlResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "RemoveTagByUrl", "imm", "")
	resp = &RemoveTagByUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveTagByUrlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
