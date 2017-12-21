package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileTagRequest struct {
	Idtype     int64  `position:"Query" name:"Idtype"`
	BeginDate  string `position:"Query" name:"BeginDate"`
	EndDate    string `position:"Query" name:"EndDate"`
	AppType    int    `position:"Query" name:"AppType"`
	Tag        string `position:"Query" name:"Tag"`
	Agsid      int64  `position:"Query" name:"Agsid"`
	AreaNumber int    `position:"Query" name:"AreaNumber"`
}

func (r ProfileTagRequest) Invoke(client *sdk.Client) (response *ProfileTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTag", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileTagResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
