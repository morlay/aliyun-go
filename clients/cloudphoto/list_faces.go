package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Faces      ListFacesFaceList
}

type ListFacesFace struct {
	Id          int64
	Name        string
	PhotosCount int
	State       string
	IsMe        bool
	Ctime       int64
	Mtime       int64
	Axis        ListFacesAxiList
	Cover       ListFacesCover
}

type ListFacesCover struct {
	Id      int64
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
	Remark  string
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

type ListFacesAxiList []string

func (list *ListFacesAxiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
