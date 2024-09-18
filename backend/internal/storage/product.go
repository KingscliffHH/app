package storage

import (
	"context"
	"time"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectStorage interface {
	GetProject(hex string) (*data.Project, error)
	GetProjects(filter bson.M) ([]*data.Project, error)
	CreateProject(project *data.Project) (*data.Project, error)
	UpdateProject(hex string, project *data.Project) (*data.Project, error)
	UpdateMetrics(hex string, metrics *data.Metrics) (*data.Metrics, error)
	MarkCompleted(hex string, date primitive.DateTime) (*data.Project, error)
	DeleteProject(hex string) error
}

type MongoProjectStorage struct {
	db          *mongo.Database
	userStorage UserStorage
}

func NewProjectStorage(db *mongo.Database, userStorage UserStorage) *MongoProjectStorage {
	return &MongoProjectStorage{db: db, userStorage: userStorage}
}

func (p *MongoProjectStorage) GetProject(hex string) (*data.Project, error) {
	users, err := p.userStorage.GetAll()
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	project := &data.Project{}
	err = p.db.Collection("projects").FindOne(context.TODO(), bson.M{"_id": id}).Decode(project)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == project.ClientRepresentative.ID {
			project.ClientRepresentative.FullName = user.FullName
			project.ClientRepresentative.ClientRole = user.ClientRole
		} else if user.ID == project.Team.ProjectLead.ID {
			project.Team.ProjectLead.FullName = user.FullName
			project.Team.ProjectLead.Email = user.Email
			project.Team.ProjectLead.Avatar = user.Avatar
			project.Team.ProjectLead.Bio = user.Bio
		} else {
			for i, member := range project.Team.TeamMembers {
				if user.ID == member.ID {
					project.Team.TeamMembers[i].FullName = user.FullName
					project.Team.TeamMembers[i].Email = user.Email
					project.Team.TeamMembers[i].Avatar = user.Avatar
					project.Team.TeamMembers[i].Bio = user.Bio
				}
			}
		}
	}

	// data.team.projectLead = getUserInfo(data.team.projectLead.id);
	// data.team.teamMembers = data.team.teamMembers.map((m) => {
	//   return getUserInfo(m.id);
	// });
	// data.clientRepresentative = getUserInfo(data.clientRepresentative.id);

	return project, nil
}

func (p *MongoProjectStorage) GetProjects(filter bson.M) ([]*data.Project, error) {
	users, err := p.userStorage.GetAll()
	if err != nil {
		return nil, err
	}

	cursor, err := p.db.Collection("projects").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	projects := []*data.Project{}
	for cursor.Next(context.TODO()) {
		project := &data.Project{}
		err := cursor.Decode(project)
		if err != nil {
			return nil, err
		}

		// "expired" projects shouldn't be shown
		if project.Status == "completed" {
			availableUntilDate := project.CompletionDate.Time().AddDate(0, 0, project.Scope.RemainsAccessibleForNDays)
			if time.Now().After(availableUntilDate) {
				continue
			}
		}

		// find user by id for project lead
		for _, user := range users {
			if user.ID == project.Team.ProjectLead.ID {
				project.Team.ProjectLead.FullName = user.FullName
				break
			}
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// CreateProject creates a new project and returns the new project
func (p *MongoProjectStorage) CreateProject(project *data.Project) (*data.Project, error) {
	res, err := p.db.Collection("projects").InsertOne(context.TODO(), project)
	if err != nil {
		return nil, err
	}

	project.ID = res.InsertedID.(primitive.ObjectID)
	return project, nil
}

func (p *MongoProjectStorage) UpdateProject(hex string, project *data.Project) (*data.Project, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	// ignore the id field
	_, err = p.db.Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": project})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *MongoProjectStorage) UpdateMetrics(hex string, metrics *data.Metrics) (*data.Metrics, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	_, err = p.db.Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"metrics": metrics}})
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func (p *MongoProjectStorage) MarkCompleted(hex string, date primitive.DateTime) (*data.Project, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	_, err = p.db.Collection("projects").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "completed", "completiondate": date}})

	if err != nil {
		return nil, err
	}

	return p.GetProject(hex)
}

func (p *MongoProjectStorage) DeleteProject(hex string) error {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}

	_, err = p.db.Collection("projects").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
