package postgres

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tsuki42/graphql-meetup/models"
)

type MeetupRepo struct {
	DB *gorm.DB
}

func (m *MeetupRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Table("MEETUP").Find(&meetups).Error
	if err != nil {
		return nil, err
	} else {
		return meetups, nil
	}
}

func (m *MeetupRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	meetup.ID = uuid.New().String()
	if err := m.DB.Table("MEETUP").Create(meetup).Error; err != nil {
		return nil, err
	} else {
		return meetup, nil
	}
}
