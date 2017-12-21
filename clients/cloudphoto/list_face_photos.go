package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFacePhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (r ListFacePhotosRequest) Invoke(client *sdk.Client) (response *ListFacePhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListFacePhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListFacePhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListFacePhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListFacePhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListFacePhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListFacePhotosResultList
}

type ListFacePhotosResult struct {
	PhotoId int64
	State   string
}

type ListFacePhotosResultList []ListFacePhotosResult

func (list *ListFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacePhotosResult)
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
