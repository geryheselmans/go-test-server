package repository

import "github.com/geryheselmans/go-test-server/model"

type InMemoryAuthorRepository struct {
	storage map[string]*model.Author
}

func NewInMemoryAuthorRepository() InMemoryAuthorRepository {
	return InMemoryAuthorRepository{
		storage: make(map[string]*model.Author),
	}
}

func (repo InMemoryAuthorRepository) Save(authorName string, author *model.Author) error {
	repo.storage[authorName] = author

	return nil
}

func (repo InMemoryAuthorRepository) FindAll() (authors []*model.Author, err error) {
	authorList := make([]*model.Author, len(repo.storage))
	i := 0
	for _, value := range repo.storage {
		authorList[i] = value
		i++
	}

	return authorList, nil
}

func (repo InMemoryAuthorRepository) FindByAuthorName(authorName string) (author *model.Author, err error) {
	return repo.storage[authorName], nil
}

func (repo InMemoryAuthorRepository) Delete(authorName string) error {
	delete(repo.storage, authorName)

	return nil
}

func (repo InMemoryAuthorRepository) Clear() error {
	for key := range repo.storage {
		delete(repo.storage, key)
	}

	return nil
}
