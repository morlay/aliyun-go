package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListSlsLogstoreInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListSlsLogstoreInfoRequest) Invoke(client *sdk.Client) (resp *ListSlsLogstoreInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListSlsLogstoreInfo", "", "")
	resp = &ListSlsLogstoreInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSlsLogstoreInfoResponse struct {
	responses.BaseResponse
	RequestId           string
	SlsLogstoreInfoList ListSlsLogstoreInfoSlsLogstoreInfoList
}

type ListSlsLogstoreInfoSlsLogstoreInfo struct {
	Id            int64
	ServiceName   string
	ComponentName string
	LogstoreName  string
	LogType       string
}

type ListSlsLogstoreInfoSlsLogstoreInfoList []ListSlsLogstoreInfoSlsLogstoreInfo

func (list *ListSlsLogstoreInfoSlsLogstoreInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSlsLogstoreInfoSlsLogstoreInfo)
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
