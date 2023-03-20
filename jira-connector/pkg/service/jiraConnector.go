package service

import "context"

type JiraConnector interface {
	Get(ctx context.Context) (string, error)
}
