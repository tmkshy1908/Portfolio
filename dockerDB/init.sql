--ユーザーの作成
-- CREATE USER docker;
--DBの作成
CREATE DATABASE linebot;
--ユーザーにDBの権限をまとめて付与
-- GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
--ユーザーを切り替え
\c linebot
--テーブルを作成
CREATE TABLE schedule(
    day timestamp without time zone primary key
);

CREATE TABLE contents(
    contents_day date not null,
    location varchar(30),
    event_title varchar(80),
    act varchar(200),
    other_info varchar(100),
    foreign key (contents_day) references schedule(day)
);

CREATE TABLE users(
    user_id text primary key,
    condition integer
);



--テーブルにデータを挿入
-- INSERT INTO book VALUES (1,'The Very Hungry Caterpillar');