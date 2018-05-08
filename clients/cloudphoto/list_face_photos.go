package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFacePhotosRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListFacePhotosRequest) Invoke(client *sdk.Client) (resp *ListFacePhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListFacePhotos", "cloudphoto", "")
	resp = &ListFacePhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFacePhotosResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Results    ListFacePhotosResultList
}

type ListFacePhotosResult struct {
	PhotoId    common.Long
	PhotoIdStr common.String
	Mtime      common.Long
	State      common.String
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
