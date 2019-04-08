-- +migrate Up

CREATE TABLE IF NOT EXISTS user(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  open_id VARCHAR(40) DEFAULT '' NOT NULL  COMMENT '用户OpenId',
  talk_id bigint not null default 0 unique comment 'talk id',
  username VARCHAR(40) DEFAULT '' NOT NULL COMMENT '用户名',
  zone VARCHAR(10) DEFAULT '' NOT NULL COMMENT '区号',
  mobile VARCHAR(20) DEFAULT '' NOT NULL COMMENT '手机号',
  nickname VARCHAR(40) DEFAULT '' NOT NULL COMMENT '昵称',
  sex SMALLINT DEFAULT 0 NOT NULL COMMENT '性别',
  password VARCHAR(100) DEFAULT '' NOT NULL  COMMENT '密码',
  email VARCHAR(40) DEFAULT '' NOT NULL  COMMENT '邮箱',
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间戳'
) CHARACTER SET utf8mb4;

insert into user( open_id, username, nickname, password, talk_id) values('232334','test','test','1',12334);