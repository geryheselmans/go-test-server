package models

type Author struct {
	authorName string
	firstName  string
	lastName   string
	mail       string
}

type AuthorRepository interface {
	Save(author *Author) error
	FindByAuthorName(authorName string) (author *Author, err error)
	Delete(authorName string) error
	Clear() error
}

func NewAuthor(authorName string, firstName string, lastName string, mail string) *Author {
	return &Author{authorName, firstName, lastName, mail}
}

func (author *Author) FindByAuthorName(authorStorage AuthorRepository, authorName string) error {
	foundAuthor, err := authorStorage.FindByAuthorName(authorName)

	if err == nil {
		return err
	}

	author.authorName = foundAuthor.authorName
	author.firstName = foundAuthor.firstName
	author.lastName = foundAuthor.lastName
	author.mail = author.mail

	return nil
}

func (author *Author) Save(authorStorage AuthorRepository) error {
	return authorStorage.Save(author)
}

func (author *Author) Delete(authorStorage AuthorRepository, authorName string) error {
	return authorStorage.Delete(authorName)
}
