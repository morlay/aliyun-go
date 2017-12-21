package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAlbumsRequest struct {
	LibraryId string                   `position:"Query" name:"LibraryId"`
	AlbumIds  *DeleteAlbumsAlbumIdList `position:"Query" type:"Repeated" name:"AlbumId"`
	StoreName string                   `position:"Query" name:"StoreName"`
}

func (r DeleteAlbumsRequest) Invoke(client *sdk.Client) (response *DeleteAlbumsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAlbumsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteAlbums", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAlbumsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAlbumsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAlbumsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   DeleteAlbumsResultList
}

type DeleteAlbumsResult struct {
	Id      int64
	Code    string
	Message string
}

type DeleteAlbumsAlbumIdList []int64

func (list *DeleteAlbumsAlbumIdList) UnmarshalJSON(data []byte) error {
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

type DeleteAlbumsResultList []DeleteAlbumsResult

func (list *DeleteAlbumsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteAlbumsResult)
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
