package redditclone

import "github.com/google/uuid"

type Thread struct {
	id          uuid.UUID `objectbox:"id"`
	title       string    `db:"title"`
	description string    `db:"description"`
}

type Post struct {
	id       uuid.UUID `db:"id"`
	threadID uuid.UUID `db:"threadID"`
	title    string    `db:"title"`
	content  string    `db:"content"`
	votes    int       `db:"votes"`
}

type Comment struct {
	id      uuid.UUID `db:"id"`
	postID  uuid.UUID `db:"postID"`
	content string    `db:"content"`
	votes   int       `db:"votes"`
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	Comments(postID uuid.UUID) ([]Comment, error)
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
