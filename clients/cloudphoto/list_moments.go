package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMomentsRequest struct {
	requests.RpcRequest
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int    `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
}

func (req *ListMomentsRequest) Invoke(client *sdk.Client) (resp *ListMomentsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListMoments", "cloudphoto", "")
	resp = &ListMomentsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMomentsResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	NextCursor string
	TotalCount int
	RequestId  string
	Action     string
	Moments    ListMomentsMomentList
}

type ListMomentsMoment struct {
	Id           int64
	LocationName string
	PhotosCount  int
	State        string
	TakenAt      int64
	Ctime        int64
	Mtime        int64
}

type ListMomentsMomentList []ListMomentsMoment

func (list *ListMomentsMomentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentsMoment)
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
