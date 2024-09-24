package services

import "github.com/plutov/gitprint/api/pkg/git"

// "github.com/plutov/gitprint/api/pkg/storage"

type Services struct {
	// Storage storage.Interface
	GithubAuth *git.Auth
}

func InitServices() (Services, error) {
	svc := Services{
		GithubAuth: git.NewAuth(),
	}

	// svc.Storage = new(storage.Sqlite)
	// if err := svc.Storage.Init(); err != nil {
	// 	return svc, fmt.Errorf("unable to init db %w", err)
	// }

	return svc, nil
}
