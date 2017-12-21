package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterTokenRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r CreateClusterTokenRequest) Invoke(client *sdk.Client) (response *CreateClusterTokenResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateClusterTokenRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "CreateClusterToken", "/clusters/[ClusterId]/token", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		CreateClusterTokenResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateClusterTokenResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateClusterTokenResponse struct {
}
