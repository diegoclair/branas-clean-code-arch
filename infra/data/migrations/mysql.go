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
		{
			Version:     2,
			Description: "Create tab_order",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_order (
					id INT NOT NULL AUTO_INCREMENT,
					code VARCHAR(12) NOT NULL,
					cpf VARCHAR(11) NOT NULL,
					coupon_id INT NULL,
					issue_date TIMESTAMP NOT NULL,
					freight DECIMAL(7,2) NOT NULL,
					sequence INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)

					INDEX fk_tab_order_tab_coupon1_idx (coupon_id ASC) VISIBLE,
					CONSTRAINT fk_tab_order_tab_coupon1
					FOREIGN KEY (coupon_id)
					REFERENCES tab_coupon (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     3,
			Description: "Create tab_order_item",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_order_item (
					order_id INT NOT NULL,
					item_id INT NOT NULL,
					price DECIMAL(7,2) NOT NULL,
					quantity INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					
					PRIMARY KEY (order_id, item_id),
					
					INDEX fk_tab_order_item_tab_order1_idx (order_id ASC) VISIBLE,
					CONSTRAINT fk_tab_order_item_tab_order1
					FOREIGN KEY (order_id)
					REFERENCES tab_order (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
					
					INDEX fk_tab_order_item_tab_item1_idx (item_id ASC) VISIBLE,
					CONSTRAINT fk_tab_order_item_tab_item1
					FOREIGN KEY (item_id)
					REFERENCES tab_item (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     4,
			Description: "Create tab_coupon",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_coupon (
					id INT NOT NULL AUTO_INCREMENT,
					code VARCHAR(50) NOT NULL,
					percentage DECIMAL(7,2) NOT NULL,
					expiration_date TIMESTAMP NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					
					PRIMARY KEY code),
					UNIQUE INDEX id_UNIQUE (code ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
