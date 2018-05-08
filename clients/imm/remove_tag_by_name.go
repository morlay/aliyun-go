package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveTagByNameRequest struct {
	requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *RemoveTagByNameRequest) Invoke(client *sdk.Client) (resp *RemoveTagByNameResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "RemoveTagByName", "imm", "")
	resp = &RemoveTagByNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveTagByNameResponse struct {
	responses.BaseResponse
	RequestId common.String
}
