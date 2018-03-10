package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTagPhotosRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	TagId     int64  `position:"Query" name:"TagId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListTagPhotosRequest) Invoke(client *sdk.Client) (resp *ListTagPhotosResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTagPhotos", "cloudphoto", "")
	resp = &ListTagPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagPhotosResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Results    ListTagPhotosResultList
}

type ListTagPhotosResult struct {
	PhotoId    int64
	PhotoIdStr string
	State      string
}

type ListTagPhotosResultList []ListTagPhotosResult

func (list *ListTagPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagPhotosResult)
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
