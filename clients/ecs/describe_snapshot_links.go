package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnapshotLinksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	SnapshotLinkIds      string `position:"Query" name:"SnapshotLinkIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeSnapshotLinksRequest) Invoke(client *sdk.Client) (response *DescribeSnapshotLinksResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSnapshotLinksRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotLinks", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSnapshotLinksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSnapshotLinksResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSnapshotLinksResponse struct {
	RequestId     string
	TotalCount    int
	PageNumber    int
	PageSize      int
	SnapshotLinks DescribeSnapshotLinksSnapshotLinkList
}

type DescribeSnapshotLinksSnapshotLink struct {
	SnapshotLinkId string
	RegionId       string
	InstanceId     string
	InstanceName   string
	SourceDiskId   string
	SourceDiskSize int
	SourceDiskType string
	TotalSize      int
	TotalCount     int
}

type DescribeSnapshotLinksSnapshotLinkList []DescribeSnapshotLinksSnapshotLink

func (list *DescribeSnapshotLinksSnapshotLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotLinksSnapshotLink)
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
