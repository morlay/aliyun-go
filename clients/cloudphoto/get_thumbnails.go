package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetThumbnailsRequest struct {
	requests.RpcRequest
	LibraryId string                    `position:"Query" name:"LibraryId"`
	PhotoIds  *GetThumbnailsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                    `position:"Query" name:"StoreName"`
	ZoomType  string                    `position:"Query" name:"ZoomType"`
}

func (req *GetThumbnailsRequest) Invoke(client *sdk.Client) (resp *GetThumbnailsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetThumbnails", "cloudphoto", "")
	resp = &GetThumbnailsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetThumbnailsResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   GetThumbnailsResultList
}

type GetThumbnailsResult struct {
	Code         common.String
	Message      common.String
	PhotoId      common.Long
	PhotoIdStr   common.String
	ThumbnailUrl common.String
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
