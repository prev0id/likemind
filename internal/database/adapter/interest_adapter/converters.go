package interest_adapter

import (
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func repoInterestToDomain(groupID int64, interests []model.Interest, selected map[int64]struct{}) []domain.Interest {
	result := make([]domain.Interest, 0, len(interests))
	for _, interest := range interests {
		if interest.GroupID != groupID {
			continue
		}

		_, isSelected := selected[interest.ID]

		result = append(result, domain.Interest{
			ID:          interest.ID,
			Name:        interest.Name,
			Description: interest.Description,
			GroupID:     interest.GroupID,
			Selected:    isSelected,
		})
	}
	return result
}

func repoUserInterestsToDomain(userInterests []model.UserInterest, groups []model.InterestGroup, interests []model.Interest) domain.Interests {
	selected := make(map[int64]struct{}, len(userInterests))
	for _, interest := range userInterests {
		selected[interest.InterestID] = struct{}{}
	}

	result := make(domain.Interests, 0, len(groups))
	for _, group := range groups {
		result = append(result, domain.InterestGroup{
			ID:        group.ID,
			Name:      group.Name,
			Interests: repoInterestToDomain(group.ID, interests, selected),
		})
	}
	return result
}

func repoGroupInterestsToDomain(groupInterests []model.GroupInterest, groups []model.InterestGroup, interests []model.Interest) domain.Interests {
	selected := make(map[int64]struct{}, len(groupInterests))
	for _, interest := range groupInterests {
		selected[interest.InterestID] = struct{}{}
	}

	result := make(domain.Interests, 0, len(groups))
	for _, group := range groups {
		result = append(result, domain.InterestGroup{
			ID:        group.ID,
			Name:      group.Name,
			Interests: repoInterestToDomain(group.ID, interests, selected),
		})
	}
	return result
}
