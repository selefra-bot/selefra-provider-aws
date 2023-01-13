



package appstream





import (




	"testing"





	"github.com/aws/aws-sdk-go-v2/service/appstream"




	"github.com/aws/aws-sdk-go-v2/service/appstream/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-aws/table_schema_generator"




)



func buildAppstreamUsageReportSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAppstreamClient(ctrl)


	object := types.UsageReportSubscription{}
	err := faker.FakeObject(&object)
	if err != nil {




		t.Fatal(err)


	}





	m.EXPECT().DescribeUsageReportSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&appstream.DescribeUsageReportSubscriptionsOutput{
			UsageReportSubscriptions: []types.UsageReportSubscription{object},


		}, nil)



	tagsOutput := appstream.ListTagsForResourceOutput{}


	err = faker.FakeObject(&tagsOutput)


	if err != nil {


		t.Fatal(err)


	}




	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsOutput, nil).AnyTimes()







	return aws_client.AwsServices{


		Appstream: m,
	}


}

func TestAppstreamUsageReportSubscriptions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppstreamUsageReportSubscriptionsGenerator{}), buildAppstreamUsageReportSubscriptionsMock, aws_client.TestOptions{})
}




