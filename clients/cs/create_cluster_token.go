package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterTokenRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *CreateClusterTokenRequest) Invoke(client *sdk.Client) (resp *CreateClusterTokenResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CreateClusterToken", "/clusters/[ClusterId]/token", "", "")
	req.Method = "POST"

	resp = &CreateClusterTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterTokenResponse struct {
	responses.BaseResponse
}
