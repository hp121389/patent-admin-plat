package version

import (
	"github.com/go-admin-team/go-admin-core/sdk/config"
	modelsAdmin "go-admin/app/admin-agent/model"
	modelsSys "go-admin/app/admin/models"
	modelsUser "go-admin/app/user-agent/models"
	"go-admin/cmd/migrate/migration/models"
	"runtime"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		err := tx.Migrator().AutoMigrate(
			new(modelsSys.SysLoginLog),
			new(modelsSys.SysUser),
			new(modelsSys.SysRole),
			new(modelsSys.SysDept),
			new(modelsSys.DeptRelation),
			new(modelsUser.Patent),
			new(modelsUser.Package),
			new(modelsUser.UserPatent),
			new(modelsUser.PatentPackage),
			new(modelsUser.StoredQuery),
			new(modelsUser.TraceLog),
			new(modelsAdmin.Report),
			new(modelsAdmin.ReportRelation),
			new(modelsAdmin.Ticket),
		)
		if err != nil {
			return err
		}
		if err := models.InitDb(tx); err != nil {
			return err
		}
		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
