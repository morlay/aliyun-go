package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FetchAlbumTagPhotosRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	TagId     int64  `position:"Query" name:"TagId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
}

func (req *FetchAlbumTagPhotosRequest) Invoke(client *sdk.Client) (resp *FetchAlbumTagPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchAlbumTagPhotos", "cloudphoto", "")
	resp = &FetchAlbumTagPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type FetchAlbumTagPhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	TotalCount int
	RequestId  string
	Action     string
	Results    FetchAlbumTagPhotosResultList
}

type FetchAlbumTagPhotosResult struct {
	PhotoId    int64
	PhotoIdStr string
	Mtime      int64
	State      string
}

type FetchAlbumTagPhotosResultList []FetchAlbumTagPhotosResult

func (list *FetchAlbumTagPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchAlbumTagPhotosResult)
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
