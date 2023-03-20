package service

import (
	"github.com/kuhotkin/jira-connector/db"
)

type jiraConnector struct {
	database db.Database
}

func New(db db.Database) JiraConnector {
	return &jiraConnector{
		db,
	}
}
