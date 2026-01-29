package app

type CommentsStorage interface {
	CreateComment(comment Comment) error
	GetComments(site, resource string) ([]Comment, error)
}
