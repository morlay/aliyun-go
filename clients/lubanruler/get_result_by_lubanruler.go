package lubanruler

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetResultByLubanrulerRequest struct {
	requests.RpcRequest
	AonePipelineId int    `position:"Query" name:"AonePipelineId"`
	Token          string `position:"Query" name:"Token"`
}

func (req *GetResultByLubanrulerRequest) Invoke(client *sdk.Client) (resp *GetResultByLubanrulerResponse, err error) {
	req.InitWithApiInfo("Lubanruler", "2017-12-28", "GetResultByLubanruler", "", "")
	resp = &GetResultByLubanrulerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetResultByLubanrulerResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RunStatus common.String
	RequestId common.String
	Data      GetResultByLubanrulerData
}

type GetResultByLubanrulerData struct {
	AonePipelineId         common.Integer
	AppName                common.String
	ScmUrl                 common.String
	ScmBranch              common.String
	TaskStatus             common.String
	BlockerCnt             common.Integer
	CriticalCnt            common.Integer
	InfoCnt                common.Integer
	MajorCnt               common.Integer
	MinorCnt               common.Integer
	Complexity             common.String
	DuplicatedLinesDensity common.String
	ReliabilityRating      common.String
	SecurityRating         common.String
	SqaleRating            common.String
	LineOfCode             common.String
	Functions              common.String
	ResultUrl              common.String
	CommentLinesDensity    common.String
}
