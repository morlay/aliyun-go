package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNoteRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeNoteRequest) Invoke(client *sdk.Client) (resp *DescribeNoteResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeNote", "", "")
	resp = &DescribeNoteResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNoteResponse struct {
	responses.BaseResponse
	RequestId  string
	Id         string
	Name       string
	ClusterId  string
	Type       string
	Paragraphs string
}
