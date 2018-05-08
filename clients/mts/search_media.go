package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchMediaRequest struct {
	requests.RpcRequest
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

func (req *SearchMediaRequest) Invoke(client *sdk.Client) (resp *SearchMediaResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchMedia", "mts", "")
	resp = &SearchMediaResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchMediaResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalNum   common.Long
	PageNumber common.Long
	PageSize   common.Long
	MediaList  SearchMediaMediaList
}

type SearchMediaMedia struct {
	MediaId      common.String
	Title        common.String
	Description  common.String
	CoverURL     common.String
	CateId       common.Long
	Duration     common.String
	Format       common.String
	Size         common.String
	Bitrate      common.String
	Width        common.String
	Height       common.String
	Fps          common.String
	PublishState common.String
	CreationTime common.String
	Tags         SearchMediaTagList
	RunIdList    SearchMediaRunIdListList
	File         SearchMediaFile
}

type SearchMediaFile struct {
	URL   common.String
	State common.String
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

type SearchMediaTagList []common.String

func (list *SearchMediaTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type SearchMediaRunIdListList []common.String

func (list *SearchMediaRunIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
