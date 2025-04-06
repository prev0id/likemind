package group_adapter

import (
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func groupDomainToModel(g domain.Group) model.Group {
	return model.Group{
		ID:          int64(g.ID),
		Name:        g.Name,
		Description: g.Description,
		AuthorID:    int64(g.Author),
	}
}

func groupModelToDomain(m model.Group) domain.Group {
	return domain.Group{
		ID:          domain.GroupID(m.ID),
		Name:        m.Name,
		Description: m.Description,
		Author:      domain.UserID(m.AuthorID),
	}
}

func postDomainToModel(p domain.Post) model.Post {
	return model.Post{
		ID:       int64(p.ID),
		GroupID:  int64(p.Group),
		Content:  p.Content,
		AuthorID: int64(p.Author),
	}
}

func postModelToDomain(m model.Post) domain.Post {
	return domain.Post{
		ID:        domain.PostID(m.ID),
		Group:     domain.GroupID(m.GroupID),
		Content:   m.Content,
		Author:    domain.UserID(m.AuthorID),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func commentDomainToModel(c domain.Comment) model.Comment {
	return model.Comment{
		ID:       int64(c.ID),
		PostID:   int64(c.PostID),
		Content:  c.Content,
		AuthorID: int64(c.Author),
	}
}

func commentModelToDomain(m model.Comment) domain.Comment {
	return domain.Comment{
		ID:        domain.CommentID(m.ID),
		PostID:    domain.PostID(m.PostID),
		Content:   m.Content,
		Author:    domain.UserID(m.AuthorID),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
