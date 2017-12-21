package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMomentPhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Direction string `position:"Query" name:"Direction"`
}

func (r ListMomentPhotosRequest) Invoke(client *sdk.Client) (response *ListMomentPhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListMomentPhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListMomentPhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListMomentPhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListMomentPhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListMomentPhotosResponse struct {
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListMomentPhotosResultList
}

type ListMomentPhotosResult struct {
	PhotoId int64
	State   string
}

type ListMomentPhotosResultList []ListMomentPhotosResult

func (list *ListMomentPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentPhotosResult)
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
