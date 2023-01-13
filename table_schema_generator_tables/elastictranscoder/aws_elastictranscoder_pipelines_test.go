

package elastictranscoder

import (




	"testing"





	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"


	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"


)









func buildElastictranscoderPipelinesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	m := mocks.NewMockElastictranscoderClient(ctrl)



	pipeline := types.Pipeline{}


	if err := faker.FakeObject(&pipeline); err != nil {


		t.Fatal(err)




	}





	m.EXPECT().ListPipelines(gomock.Any(), gomock.Any()).AnyTimes().Return(


		&elastictranscoder.ListPipelinesOutput{Pipelines: []types.Pipeline{pipeline}},
		nil,




	)





	job := types.Job{}




	if err := faker.FakeObject(&job); err != nil {


		t.Fatal(err)


	}




	job.PipelineId = pipeline.Id



	m.EXPECT().ListJobsByPipeline(gomock.Any(), gomock.Any()).AnyTimes().Return(


		&elastictranscoder.ListJobsByPipelineOutput{Jobs: []types.Job{job}},




		nil,




	)





	return aws_client.AwsServices{




		Elastictranscoder: m,


	}
}

func TestElastictranscoderPipelines(t *testing.T) {


	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElastictranscoderPipelinesGenerator{}), buildElastictranscoderPipelinesMock, aws_client.TestOptions{})


}


