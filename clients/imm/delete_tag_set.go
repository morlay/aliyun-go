package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTagSetRequest struct {
	requests.RpcRequest
	LazyMode   string `position:"Query" name:"LazyMode"`
	Project    string `position:"Query" name:"Project"`
	SetId      string `position:"Query" name:"SetId"`
	CheckEmpty string `position:"Query" name:"CheckEmpty"`
}

func (req *DeleteTagSetRequest) Invoke(client *sdk.Client) (resp *DeleteTagSetResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteTagSet", "imm", "")
	resp = &DeleteTagSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTagSetResponse struct {
	responses.BaseResponse
	RequestId common.String
}
