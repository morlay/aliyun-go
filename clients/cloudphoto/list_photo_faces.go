package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhotoFacesRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
}

func (r ListPhotoFacesRequest) Invoke(client *sdk.Client) (response *ListPhotoFacesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPhotoFacesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoFaces", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListPhotoFacesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPhotoFacesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPhotoFacesResponse struct {
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
