package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeParagraphRequest) Invoke(client *sdk.Client) (resp *DescribeParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeParagraph", "", "")
	resp = &DescribeParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeParagraphResponse struct {
	responses.BaseResponse
	RequestId string
	Paragraph string
}
