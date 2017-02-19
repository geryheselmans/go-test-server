package models

type Author struct {
	AuthorName string `json:"authorName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Mail       string `json:"mail"`
}

type AuthorRepository interface {
	Save(authorName string, author *Author) error
	FindAll() (authors []*Author, err error)
	FindByAuthorName(authorName string) (author *Author, err error)
	Delete(authorName string) error
	Clear() error
}

func (author *Author) FindByAuthorName(authorStorage AuthorRepository, authorName string) error {
	foundAuthor, err := authorStorage.FindByAuthorName(authorName)

	if err == nil {
		return err
	}

	author.AuthorName = foundAuthor.AuthorName
	author.FirstName = foundAuthor.FirstName
	author.LastName = foundAuthor.LastName
	author.Mail = author.Mail

	return nil
}

func (author *Author) Save(authorStorage AuthorRepository) error {
	return authorStorage.Save(author.AuthorName, author)
}

func (author *Author) Delete(authorStorage AuthorRepository, authorName string) error {
	return authorStorage.Delete(authorName)
}
