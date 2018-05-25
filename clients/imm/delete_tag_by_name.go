package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTagByNameRequest struct {
	requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *DeleteTagByNameRequest) Invoke(client *sdk.Client) (resp *DeleteTagByNameResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteTagByName", "imm", "")
	resp = &DeleteTagByNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTagByNameResponse struct {
	responses.BaseResponse
	RequestId common.String
}
