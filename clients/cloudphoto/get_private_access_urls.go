package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPrivateAccessUrlsRequest struct {
	LibraryId string                           `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPrivateAccessUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                           `position:"Query" name:"StoreName"`
	ZoomType  string                           `position:"Query" name:"ZoomType"`
}

func (r GetPrivateAccessUrlsRequest) Invoke(client *sdk.Client) (response *GetPrivateAccessUrlsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPrivateAccessUrlsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPrivateAccessUrls", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetPrivateAccessUrlsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPrivateAccessUrlsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPrivateAccessUrlsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetPrivateAccessUrlsResultList
}

type GetPrivateAccessUrlsResult struct {
	Code      string
	Message   string
	PhotoId   int64
	AccessUrl string
}

type GetPrivateAccessUrlsPhotoIdList []int64

func (list *GetPrivateAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type GetPrivateAccessUrlsResultList []GetPrivateAccessUrlsResult

func (list *GetPrivateAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPrivateAccessUrlsResult)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
