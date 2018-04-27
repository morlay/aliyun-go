package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetSoftwarePasswordRequest struct {
	requests.RpcRequest
	Password  string `position:"Query" name:"Password"`
	ClusterId string `position:"Query" name:"ClusterId"`
	Username  string `position:"Query" name:"Username"`
}

func (req *ResetSoftwarePasswordRequest) Invoke(client *sdk.Client) (resp *ResetSoftwarePasswordResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResetSoftwarePassword", "", "")
	resp = &ResetSoftwarePasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetSoftwarePasswordResponse struct {
	responses.BaseResponse
	RequestId string
	Success   string
	ErrCode   string
	ErrMsg    string
}
