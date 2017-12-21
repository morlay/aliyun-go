package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterNameRequest struct {
}

func (r ModifyClusterNameRequest) Invoke(client *sdk.Client) (response *ModifyClusterNameResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ModifyClusterNameRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "ModifyClusterName", "/clusters/[ClusterId]/name/[ClusterName]", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ModifyClusterNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyClusterNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyClusterNameResponse struct {
}
