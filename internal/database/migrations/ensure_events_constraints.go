package migrations

import "gorm.io/gorm"

func EnsureEventsConstraints(db *gorm.DB) error {
	if err := db.Exec(
		"CREATE UNIQUE INDEX IF NOT EXISTS idx_events_single_root_per_repository ON events (repository_id) WHERE parent_id IS NULL",
	).Error; err != nil {
		return err
	}

	return db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1
				FROM pg_constraint
				WHERE conname = 'fk_repositories_last_event'
			) THEN
				ALTER TABLE repositories
				ADD CONSTRAINT fk_repositories_last_event
				FOREIGN KEY (last_event_id)
				REFERENCES events(id)
				ON UPDATE CASCADE
				ON DELETE SET NULL;
			END IF;
		END $$
	`).Error
}
