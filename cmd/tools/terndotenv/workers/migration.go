package workers

import (
	"fmt"
	"github.com/urfave/cli"
	"os/exec"
)

func ErrorRecovery() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func Migration(c *cli.Context) error {
	if c.IsSet("version") {
		version := fmt.Sprintf("%d", c.Int("version"))
		return migrateToVersion(version)
	} else {
		return migrateToLatest()
	}
}

func migrateToLatest() error {
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)
	return cmd.Run()
}

func migrateToVersion(version string) error {
	cmd := exec.Command(
		"tern",
		"migrate",
		"--destination",
		version,
	)
	return cmd.Run()
}

func CreateMigration(c *cli.Context) error {
	migrationName := c.String("name")
	cmd := exec.Command(
		"tern",
		"new",
		"--migrations",
		"internal/store/pgstore/migrations",
		migrationName,
	)
	fmt.Println(cmd.String())
	return cmd.Run()
}
