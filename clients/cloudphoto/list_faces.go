package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFacesRequest struct {
	requests.RpcRequest
	Cursor      string `position:"Query" name:"Cursor"`
	HasFaceName string `position:"Query" name:"HasFaceName"`
	Size        int    `position:"Query" name:"Size"`
	LibraryId   string `position:"Query" name:"LibraryId"`
	StoreName   string `position:"Query" name:"StoreName"`
	State       string `position:"Query" name:"State"`
	Direction   string `position:"Query" name:"Direction"`
}

func (req *ListFacesRequest) Invoke(client *sdk.Client) (resp *ListFacesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListFaces", "cloudphoto", "")
	resp = &ListFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFacesResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Faces      ListFacesFaceList
}

type ListFacesFace struct {
	Id          common.Long
	IdStr       common.String
	Name        common.String
	PhotosCount common.Integer
	State       common.String
	IsMe        bool
	Ctime       common.Long
	Mtime       common.Long
	Axis        ListFacesAxiList
	Cover       ListFacesCover
}

type ListFacesCover struct {
	Id      common.Long
	IdStr   common.String
	Title   common.String
	FileId  common.String
	State   common.String
	Md5     common.String
	IsVideo bool
	Width   common.Long
	Height  common.Long
	Ctime   common.Long
	Mtime   common.Long
	Remark  common.String
}

type ListFacesFaceList []ListFacesFace

func (list *ListFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacesFace)
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

type ListFacesAxiList []common.String

func (list *ListFacesAxiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
