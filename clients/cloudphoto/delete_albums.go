package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteAlbumsRequest struct {
	requests.RpcRequest
	LibraryId string                   `position:"Query" name:"LibraryId"`
	AlbumIds  *DeleteAlbumsAlbumIdList `position:"Query" type:"Repeated" name:"AlbumId"`
	StoreName string                   `position:"Query" name:"StoreName"`
}

func (req *DeleteAlbumsRequest) Invoke(client *sdk.Client) (resp *DeleteAlbumsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteAlbums", "cloudphoto", "")
	resp = &DeleteAlbumsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAlbumsResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   DeleteAlbumsResultList
}

type DeleteAlbumsResult struct {
	Id      common.Long
	IdStr   common.String
	Code    common.String
	Message common.String
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
