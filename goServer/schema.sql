CREATE TABLE user (
	id serial PRIMARY KEY,
	name varchar(42) NOT NULL,
	pin int NOT NULL,
	joindate date NOT NULL default current_date
);
CREATE TABLE task(
	id serial primary key,
	name varchar(42) NOT NULL,
	discription text NOT NULL,
	created date NOT NULL default current_date,
	createdby_user_id serial not NULL,
	done boolean default false,
	CONSTRAINT fk_user FOREIGN KEY(createdby_user_id) REFERENCES user(id)
);
CREATE TABLE user_task(
	id serial primary key,
	user_id serial,
	task_id serial,
	added date NOT NULL default current_date,
	CONSTRAINT fk_user FOREIGN KEY(user_id) references user(id),
	CONSTRAINT fk_task FOREIGN KEY(task_id) REFERENCES task(id)
);