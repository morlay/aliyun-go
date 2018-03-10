package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMomentPhotosRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListMomentPhotosRequest) Invoke(client *sdk.Client) (resp *ListMomentPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListMomentPhotos", "cloudphoto", "")
	resp = &ListMomentPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMomentPhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListMomentPhotosResultList
}

type ListMomentPhotosResult struct {
	PhotoId    int64
	PhotoIdStr string
	State      string
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
