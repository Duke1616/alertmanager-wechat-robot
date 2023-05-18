package impl

import (
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dashboard"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) QueryAlertMessage(ctx context.Context, req *dashboard.Request) (*dashboard.Request, error) {
	query, err := newQueryDashboardRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query policy filter: %s", query.AggregatePipeline())

	//
	//matchPipeline := bson.D{
	//	{"$match", bson.D{
	//		{"level", bson.D{{"$eq", "P4"}}},
	//	}},
	//}
	//
	//pipeline := bson.D{
	//	{"$group", bson.D{
	//		{"_id", "$instance_name"},
	//		{"count", bson.D{{"$sum", 1}}},
	//	}},
	//}
	//
	//sortPipeline := bson.D{
	//	{"$sort", bson.D{{"count", -1}}},
	//}
	//

	resp, err := s.col.Aggregate(ctx, query.AggregatePipeline(), query.AggregateOptions())

	if err != nil {
		return nil, err
	}

	var results []bson.M

	if err = resp.All(context.TODO(), &results); err != nil {
		s.log.Error(err)
		//panic(err)
	}
	for _, result := range results {
		fmt.Printf("Average price of %v tea options: $%v \n", result["_id"], result["count"])
		fmt.Printf("Number of %v tea options: %v \n\n", result["_id"], result["count"])
	}

	return nil, nil
}
