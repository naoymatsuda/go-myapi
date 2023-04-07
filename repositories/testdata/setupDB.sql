create table if not exists articles (
	article_id INTEGER PRIMARY KEY AUTOINCREMENT,
	title STRING NOT NULL,
	contents TEXT NOT NULL,
	username STRING NOT NULL,
	nice INTEGER NOT NULL,
	created_at DATETIME
);

create table if not exists comments (
	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
	article_id INTEGER NOT NULL,
	message TEXT NOT NULL,
	created_at DATETIME,
	foreign key (article_id) references articles(article_id)
);

insert into articles (title, contents, username, nice, created_at) values
	('firstPost', 'This is my first blog', 'saki', 2, time.now());

insert into articles (title, contents, username, nice) values
	('2nd', 'Second blog post', 'saki', 4);

insert into comments (article_id, message, created_at) values
	(1, '1st comment yeah', time.now());

insert into comments (article_id, message) values
	(1, 'welcome');
