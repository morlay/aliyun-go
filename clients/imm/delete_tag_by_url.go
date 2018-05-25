package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTagByUrlRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *DeleteTagByUrlRequest) Invoke(client *sdk.Client) (resp *DeleteTagByUrlResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteTagByUrl", "imm", "")
	resp = &DeleteTagByUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTagByUrlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
