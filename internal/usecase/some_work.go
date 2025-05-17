package usecase

import (
	"context"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel/tracer"
)

func (u *UseCase) SomeWork(ctx context.Context) error {
	ctx, span := tracer.Start(ctx, "usecase SomeWork")
	defer span.End()

	// log.Info().Msg("SomeWork called")

	// Пример вызова клиента
	//p, err := u.profile.GetProfile(ctx, "8638341a-b68a-4291-84ee-94b147afeff9")
	//if err != nil {
	//	return fmt.Errorf("SomeWork: %w", err)
	//}
	//
	//fmt.Println(p)

	return nil
}
