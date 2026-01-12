package app

type Storage interface {
	CreateComment(comment Comment) error
	GetComments(site, resource string) ([]Comment, error)
}
