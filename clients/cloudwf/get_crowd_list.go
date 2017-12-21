package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCrowdListRequest struct {
	requests.RpcRequest
	Gsid      int64  `position:"Query" name:"Gsid"`
	ClassType int    `position:"Query" name:"ClassType"`
	GsType    string `position:"Query" name:"GsType"`
	EndTime   string `position:"Query" name:"EndTime"`
	Page      int    `position:"Query" name:"Page"`
	StartTime string `position:"Query" name:"StartTime"`
	Per       int    `position:"Query" name:"Per"`
}

func (req *GetCrowdListRequest) Invoke(client *sdk.Client) (resp *GetCrowdListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetCrowdList", "", "")
	resp = &GetCrowdListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetCrowdListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
