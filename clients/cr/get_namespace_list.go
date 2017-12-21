package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNamespaceListRequest struct {
}

func (r GetNamespaceListRequest) Invoke(client *sdk.Client) (response *GetNamespaceListResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetNamespaceListRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetNamespaceList", "/namespace", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetNamespaceListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetNamespaceListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetNamespaceListResponse struct {
}
