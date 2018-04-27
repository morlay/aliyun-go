package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAlbumPhotosRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListAlbumPhotosRequest) Invoke(client *sdk.Client) (resp *ListAlbumPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListAlbumPhotos", "cloudphoto", "")
	resp = &ListAlbumPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAlbumPhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListAlbumPhotosResultList
}

type ListAlbumPhotosResult struct {
	PhotoId    int64
	PhotoIdStr string
	Mtime      int64
	State      string
}

type ListAlbumPhotosResultList []ListAlbumPhotosResult

func (list *ListAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumPhotosResult)
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
