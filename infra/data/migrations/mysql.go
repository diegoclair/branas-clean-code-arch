package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
//Only work doing 1 command per version, you cannot create two tables in the same script, need to create a new version
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Create tab_item",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_item (
					id INT NOT NULL AUTO_INCREMENT,
					category VARCHAR(100) NOT NULL,
					description VARCHAR(250) NOT NULL,
					price DECIMAL(7,2) NOT NULL,
					width DECIMAL(7,2) NOT NULL,
					height DECIMAL(7,2) NOT NULL,
					length DECIMAL(7,2) NOT NULL,
					weight DECIMAL(7,2) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
