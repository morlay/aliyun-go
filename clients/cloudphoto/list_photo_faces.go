package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhotoFacesRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (req *ListPhotoFacesRequest) Invoke(client *sdk.Client) (resp *ListPhotoFacesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoFaces", "cloudphoto", "")
	resp = &ListPhotoFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhotoFacesResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Faces     ListPhotoFacesFaceList
}

type ListPhotoFacesFace struct {
	FaceId   int64
	FaceName string
	Axis     ListPhotoFacesAxiList
}

type ListPhotoFacesFaceList []ListPhotoFacesFace

func (list *ListPhotoFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoFacesFace)
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

type ListPhotoFacesAxiList []string

func (list *ListPhotoFacesAxiList) UnmarshalJSON(data []byte) error {
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
