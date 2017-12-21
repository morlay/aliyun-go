package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetThumbnailsRequest struct {
	LibraryId string                    `position:"Query" name:"LibraryId"`
	PhotoIds  *GetThumbnailsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                    `position:"Query" name:"StoreName"`
	ZoomType  string                    `position:"Query" name:"ZoomType"`
}

func (r GetThumbnailsRequest) Invoke(client *sdk.Client) (response *GetThumbnailsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetThumbnailsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetThumbnails", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetThumbnailsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetThumbnailsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetThumbnailsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetThumbnailsResultList
}

type GetThumbnailsResult struct {
	Code         string
	Message      string
	PhotoId      int64
	ThumbnailUrl string
}

type GetThumbnailsPhotoIdList []int64

func (list *GetThumbnailsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetThumbnailsResultList []GetThumbnailsResult

func (list *GetThumbnailsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThumbnailsResult)
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
