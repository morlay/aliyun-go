package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDownloadUrlsRequest struct {
	requests.RpcRequest
	LibraryId string                      `position:"Query" name:"LibraryId"`
	PhotoIds  *GetDownloadUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                      `position:"Query" name:"StoreName"`
}

func (req *GetDownloadUrlsRequest) Invoke(client *sdk.Client) (resp *GetDownloadUrlsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetDownloadUrls", "cloudphoto", "")
	resp = &GetDownloadUrlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDownloadUrlsResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetDownloadUrlsResultList
}

type GetDownloadUrlsResult struct {
	Code        string
	Message     string
	PhotoId     int64
	PhotoIdStr  string
	DownloadUrl string
}

type GetDownloadUrlsPhotoIdList []int64

func (list *GetDownloadUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetDownloadUrlsResultList []GetDownloadUrlsResult

func (list *GetDownloadUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDownloadUrlsResult)
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
