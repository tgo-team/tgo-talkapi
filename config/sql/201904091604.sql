-- +migrate Up

insert into user(open_id, username, nickname, password, talk_id)
values ('123456', 'test1', 'test1', '1', 123456);

insert into user(open_id, username, nickname, password, talk_id)
values ('234567', 'test2', 'test2', '1', 234567);

insert into user(open_id, username, nickname, password, talk_id)
values ('345678', 'test3', 'test3', '1', 345678);

insert into contacts(open_id, relation_open_id) values ('123456','234567');
insert into contacts(open_id, relation_open_id) values ('123456','345678');
