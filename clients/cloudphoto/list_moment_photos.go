package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code       common.String
	Message    common.String
	NextCursor common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Results    ListMomentPhotosResultList
}

type ListMomentPhotosResult struct {
	PhotoId    common.Long
	PhotoIdStr common.String
	State      common.String
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
