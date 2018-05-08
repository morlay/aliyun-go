package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Moments    ListMomentsMomentList
}

type ListMomentsMoment struct {
	Id           common.Long
	IdStr        common.String
	LocationName common.String
	PhotosCount  common.Integer
	State        common.String
	TakenAt      common.Long
	Ctime        common.Long
	Mtime        common.Long
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
