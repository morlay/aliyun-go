package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	CateId               string `position:"Query" name:"CateId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	From                 string `position:"Query" name:"From"`
	SortBy               string `position:"Query" name:"SortBy"`
	To                   string `position:"Query" name:"To"`
	Tag                  string `position:"Query" name:"Tag"`
	KeyWord              string `position:"Query" name:"KeyWord"`
}

func (r SearchMediaRequest) Invoke(client *sdk.Client) (response *SearchMediaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SearchMediaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchMedia", "", "")

	resp := struct {
		*responses.BaseResponse
		SearchMediaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SearchMediaResponse

	err = client.DoAction(&req, &resp)
	return
}

type SearchMediaResponse struct {
	RequestId  string
	TotalNum   int64
	PageNumber int64
	PageSize   int64
	MediaList  SearchMediaMediaList
}

type SearchMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         SearchMediaTagList
	RunIdList    SearchMediaRunIdListList
	File         SearchMediaFile
}

type SearchMediaFile struct {
	URL   string
	State string
}

type SearchMediaMediaList []SearchMediaMedia

func (list *SearchMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaMedia)
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

type SearchMediaTagList []string

func (list *SearchMediaTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type SearchMediaRunIdListList []string

func (list *SearchMediaRunIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
