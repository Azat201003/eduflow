package server

import (
	"context"

	"summary-service/server/db"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/summary"
	"gorm.io/gorm"
)

type summaryServiceServer struct {
	pb.UnimplementedSummaryServiceServer
	dbm *db.DBManger
}

func (s *summaryServiceServer) GetSummaryById(context context.Context, id *pb.Id) (*pb.Summary, error) {
	summary := db.Summary{ID: id.Id}
	err := s.dbm.FindSummary(&summary)
	return &pb.Summary{
		Id:          &pb.Id{Id: summary.ID},
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    &pb.Id{Id: summary.AuthorId},
	}, err
}

func (s *summaryServiceServer) CreateSummary(ctx context.Context, summary *pb.Summary) (*pb.Id, error) {
	db_summary := db.Summary{
		Title:       summary.Title,
		Description: summary.Description,
		FilePath:    summary.FilePath,
		AuthorId:    summary.AuthorId.Id,
	}
	err := s.dbm.CreateSummary(&db_summary)
	return &pb.Id{Id: db_summary.ID}, err
}

// func (s *summaryServiceServer) UpdateSummary(ctx context.Context, summary *pb.Summary) error {
// 	db_summary := db.Summary{
// 		Title:       summary.Title,
// 		Description: summary.Description,
// 		FilePath:    summary.FilePath,
// 		ID:          summary.Id.Id,
// 		AuthorId:    summary.AuthorId.Id,
// 	}
// 	err := db.UpdateSummary(s.db, &db_summary)
// 	return err
// }

// func (s *summaryServiceServer) DeleteSummary(ctx context.Context, id *pb.Id) error {
// 	db_summary := db.Summary{
// 		ID: id.Id,
// 	}
// 	err := s.dbm.(&db_summary)
// 	return err
// }

/*
func (s *summaryServiceServer) ListSummaries(ctx context.Context, empty *pb.Empty) (*pb.Summaries, error) {
	var summaries []db.Summary
	err := db.ListSummaries(s.db, &summaries)
	if err != nil {
		return nil, err
	}
	var pbSummaries []*pb.Summary
	for _, summary := range summaries {
		pbSummaries = append(pbSummaries, &pb.Summary{
			Id:          &pb.Id{Id: summary.ID},
			Title:       summary.Title,
			Description: summary.Description,
			FilePath:    summary.FilePath,
		})
	}
	return &pb.Summaries{Summaries: pbSummaries}, nil
}
*/

func NewServer(db_gorm *gorm.DB) pb.SummaryServiceServer {
	server := new(summaryServiceServer)
	server.dbm = &db.DBManger{DB: db_gorm}
	return server
}
