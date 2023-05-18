package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dashboard"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryDashboardRequest(req *dashboard.Request) (*queryDashboardRequest, error) {
	return &queryDashboardRequest{
		Request: req}, nil
}

type queryDashboardRequest struct {
	*dashboard.Request
}

func (r *queryDashboardRequest) AggregateOptions() *options.AggregateOptions {
	opt := &options.AggregateOptions{}

	return opt
}

func (r *queryDashboardRequest) AggregatePipeline() mongo.Pipeline {
	pipeline := mongo.Pipeline{}

	filter := bson.M{}
	filter["level"] = bson.M{"$eq": "P4"}

	matchPipeline := bson.D{
		{"$match", filter},
	}

	pipeline = append(pipeline, matchPipeline)

	pipelineGroup := bson.D{
		{"$group", bson.D{
			{"_id", "$instance_name"},
			{"count", bson.D{{"$sum", 1}}},
		}},
	}

	pipeline = append(pipeline, pipelineGroup)

	sortPipeline := bson.D{
		{"$sort", bson.D{{"count", -1}}},
	}

	pipeline = append(pipeline, sortPipeline)

	return pipeline
}
