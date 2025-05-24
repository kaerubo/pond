//go:build wireinject

package wire

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/kaerubo/kaeruashi/internal/handler"
	"github.com/kaerubo/kaeruashi/internal/infra/postgres"
	"github.com/kaerubo/kaeruashi/internal/repository"
	"github.com/kaerubo/kaeruashi/internal/usecase"
)

func InitializeHandler(db *sql.DB) *handler.KeroHandler {
	wire.Build(
		postgres.NewKeroRepository,
		wire.Bind(new(repository.KeroSaver), new(*postgres.KeroRepository)),
		wire.Bind(new(repository.KeroByIDFinder), new(*postgres.KeroRepository)),
		wire.Bind(new(repository.KeroLister), new(*postgres.KeroRepository)),
		wire.Bind(new(repository.KeroUpdater), new(*postgres.KeroRepository)),
		wire.Bind(new(repository.KeroDeleter), new(*postgres.KeroRepository)),
		usecase.NewKeroCreator,
		usecase.NewKeroByIDReader,
		usecase.NewKeroLister,
		usecase.NewKeroUpdater,
		usecase.NewKeroDeleter,
		handler.NewKeroHandler,
	)
	return nil
}
