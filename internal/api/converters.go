package api

import (
	"context"
	"likemind/internal/domain"
	"likemind/website/view"
)

func profileFromDomainToView(
	userID domain.UserID,
	user domain.User,
	contacts []domain.Contact,
	pictures []domain.PictureID,
	interests []domain.InterestGroup,
) *view.Profile {
	return &view.Profile{
		Name:        user.Name,
		Surname:     user.Surname,
		Nickname:    user.Nickname,
		About:       user.About,
		Location:    user.Location,
		DateOfBirth: user.DateOfBirth,
		Owner:       userID == user.ID,

		PictureID: convertPicturesIDs(pictures),
		Contacts:  contactsDomainToView(contacts),

		Interests: interestGroupDomainToView(interests),
	}
}

func convertPicturesIDs(picutres []domain.PictureID) string {
	if len(picutres) > 0 {
		return string(picutres[0])
	}
	return ""
}

func contactsDomainToView(contacts []domain.Contact) []view.Contact {
	result := make([]view.Contact, 0, len(contacts))
	for _, contact := range contacts {
		result = append(result, view.Contact{
			ID:       contact.ID,
			Platform: contact.Platform,
			Value:    contact.Value,
		})
	}
	return result
}

func interestGroupDomainToView(interests []domain.InterestGroup) []view.GroupedInterests {
	result := make([]view.GroupedInterests, 0, len(interests))
	for _, group := range interests {
		result = append(result, view.GroupedInterests{
			Name:      group.Name,
			Interests: interestsDomainToView(group.Interests),
		})
	}
	return result
}

func interestsDomainToView(interests []domain.Interest) []view.Interest {
	result := make([]view.Interest, 0, len(interests))
	for _, interest := range interests {
		result = append(result, view.Interest{
			ID:          int64(interest.ID),
			Name:        interest.Name,
			Description: interest.Description,
			Selected:    interest.Selected,
		})
	}
	return result
}

func (s *Server) groupDomainToView(ctx context.Context, group domain.Group, posts []domain.Post, intersts []domain.InterestGroup) (*view.Group, error) {
	convertedPosts := make([]*view.Post, 0, len(posts))
	for _, post := range posts {
		converted, err := s.postDomainToView(ctx, post)
		if err != nil {
			return nil, err
		}
		convertedPosts = append(convertedPosts, converted)
	}

	author, err := s.getProfile(ctx, group.Author)
	if err != nil {
		return nil, err
	}

	return &view.Group{
		Name:        group.Name,
		Description: group.Description,
		Author:      author,
		Posts:       convertedPosts,
		Interests:   interestGroupDomainToView(intersts),
	}, nil
}

func (s *Server) postDomainToView(ctx context.Context, post domain.Post) (*view.Post, error) {
	comments := make([]*view.Comment, 0, len(post.Comments))
	for _, comment := range post.Comments {
		converted, err := s.commentDomainToView(ctx, comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, converted)
	}

	author, err := s.getProfile(ctx, post.Author)
	if err != nil {
		return nil, err
	}

	return &view.Post{
		Author:    author,
		Comments:  comments,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (s *Server) commentDomainToView(ctx context.Context, comment domain.Comment) (*view.Comment, error) {
	author, err := s.getProfile(ctx, comment.Author)
	if err != nil {
		return nil, err
	}

	return &view.Comment{
		Content:   comment.Content,
		Author:    author,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.CreatedAt,
	}, nil
}
