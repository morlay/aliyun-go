package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetFramedPhotoUrlsRequest struct {
	requests.RpcRequest
	FrameId   string                         `position:"Query" name:"FrameId"`
	LibraryId string                         `position:"Query" name:"LibraryId"`
	PhotoIds  *GetFramedPhotoUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                         `position:"Query" name:"StoreName"`
}

func (req *GetFramedPhotoUrlsRequest) Invoke(client *sdk.Client) (resp *GetFramedPhotoUrlsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetFramedPhotoUrls", "cloudphoto", "")
	resp = &GetFramedPhotoUrlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFramedPhotoUrlsResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetFramedPhotoUrlsResultList
}

type GetFramedPhotoUrlsResult struct {
	Code           string
	Message        string
	PhotoId        int64
	PhotoIdStr     string
	FramedPhotoUrl string
}

type GetFramedPhotoUrlsPhotoIdList []int64

func (list *GetFramedPhotoUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetFramedPhotoUrlsResultList []GetFramedPhotoUrlsResult

func (list *GetFramedPhotoUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetFramedPhotoUrlsResult)
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
