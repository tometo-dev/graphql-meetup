package dataloader

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/tsuki42/graphql-meetup/models"
	"net/http"
	"time"
)

const userloaderKey = "userloader"

func DataloaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				err := db.Table("USER").Where("id in (?)", ids).Find(&users).Error
				if err != nil {
					return nil, []error{err}
				}

				/*
					Preserve the ordering of users
				*/
				userMap := make(map[string]*models.User, len(users))

				for _, user := range users {
					userMap[user.ID] = user
				}

				result := make([]*models.User, len(ids))

				for i, id := range ids {
					result[i] = userMap[id]
				}

				return result, nil

			},
		}
		ctx := context.WithValue(r.Context(), userloaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
