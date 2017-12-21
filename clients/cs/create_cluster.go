package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterRequest struct {
}

func (r CreateClusterRequest) Invoke(client *sdk.Client) (response *CreateClusterResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateClusterRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "CreateCluster", "/clusters", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		CreateClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateClusterResponse struct {
}
