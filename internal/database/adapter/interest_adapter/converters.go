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
			ID:          domain.InterestID(interest.ID),
			Name:        interest.Name,
			Description: interest.Description,
			GroupID:     interest.GroupID,
			Selected:    isSelected,
		})
	}
	return result
}

func repoUserInterestsToDomain(userInterests []model.UserInterest, groups []model.InterestGroup, interests []model.Interest) []domain.InterestGroup {
	selected := make(map[int64]struct{}, len(userInterests))
	for _, interest := range userInterests {
		selected[interest.InterestID] = struct{}{}
	}

	result := make([]domain.InterestGroup, 0, len(groups))
	for _, group := range groups {
		result = append(result, domain.InterestGroup{
			Name:      group.Name,
			Interests: repoInterestToDomain(group.ID, interests, selected),
		})
	}
	return result
}

func repoGroupInterestsToDomain(groupInterests []model.GroupInterest, groups []model.InterestGroup, interests []model.Interest) []domain.InterestGroup {
	selected := make(map[int64]struct{}, len(groupInterests))
	for _, interest := range groupInterests {
		selected[interest.InterestID] = struct{}{}
	}

	result := make([]domain.InterestGroup, 0, len(groups))
	for _, group := range groups {
		result = append(result, domain.InterestGroup{
			Name:      group.Name,
			Interests: repoInterestToDomain(group.ID, interests, selected),
		})
	}
	return result
}

func repoInterestIDs(interests []model.UserInterest) []int64 {
	result := make([]int64, 0, len(interests))
	for _, interest := range interests {
		result = append(result, int64(interest.InterestID))
	}
	return result
}

func repoUserIDsToDomain(results []model.SearchResult) []domain.UserID {
	result := make([]domain.UserID, 0, len(results))
	for _, id := range results {
		result = append(result, domain.UserID(id.ID))
	}
	return result
}

func repoGroupIDsToDomain(ids []model.SearchResult) []domain.GroupID {
	result := make([]domain.GroupID, 0, len(ids))
	for _, id := range ids {
		result = append(result, domain.GroupID(id.ID))
	}
	return result
}
