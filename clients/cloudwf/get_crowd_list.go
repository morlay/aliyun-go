package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCrowdListRequest struct {
	Gsid      int64  `position:"Query" name:"Gsid"`
	ClassType int    `position:"Query" name:"ClassType"`
	GsType    string `position:"Query" name:"GsType"`
	EndTime   string `position:"Query" name:"EndTime"`
	Page      int    `position:"Query" name:"Page"`
	StartTime string `position:"Query" name:"StartTime"`
	Per       int    `position:"Query" name:"Per"`
}

func (r GetCrowdListRequest) Invoke(client *sdk.Client) (response *GetCrowdListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetCrowdListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetCrowdList", "", "")

	resp := struct {
		*responses.BaseResponse
		GetCrowdListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetCrowdListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetCrowdListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
