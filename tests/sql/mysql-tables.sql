USE db;

DROP TABLE IF EXISTS phone, company;

CREATE TABLE company
(
	id int auto_increment primary key,
    name varchar(50) not null
);

CREATE TABLE phone
(
	id int auto_increment primary key,
	company_id int not null,
	number varchar(30), foreign key (company_id) REFERENCES company(id)
);

		INSERT INTO db.company
			(
			name
			)
		VALUES
			(
				"BBP"
	),
			(
				"ASD"
	),
			(
				"PQJ"
	),
			(
				"RYZ"
	);

		INSERT INTO db.phone
			(
			company_id,
			number
			)
		VALUES
			(
				1,
				"1184587458"
	),
			(
				2,
				"1141241233"
	),
			(
				3,
				"1115241243"
	),
			(
				4,
				"1131241242"
	);