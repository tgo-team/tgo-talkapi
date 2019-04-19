-- +migrate Up

CREATE TABLE IF NOT EXISTS user
(
  id         BIGINT PRIMARY KEY AUTO_INCREMENT,
  open_id    VARCHAR(40)  DEFAULT '' NOT NULL unique COMMENT '用户OpenId',
  talk_id    bigint                  not null default 0 unique comment 'talk id',
  username   VARCHAR(40)  DEFAULT '' NOT NULL COMMENT '用户名',
  zone       VARCHAR(10)  DEFAULT '' NOT NULL COMMENT '区号',
  mobile     VARCHAR(20)  DEFAULT '' NOT NULL COMMENT '手机号',
  nickname   VARCHAR(40)  DEFAULT '' NOT NULL COMMENT '昵称',
  sex        SMALLINT     DEFAULT 0  NOT NULL COMMENT '性别',
  password   VARCHAR(100) DEFAULT '' NOT NULL COMMENT '密码',
  email      VARCHAR(40)  DEFAULT '' NOT NULL COMMENT '邮箱',
  created_at timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间戳',
  is_deleted smallint                not null default 0 comment '是否已删除 0.否 1.是',
  index(is_deleted)
) CHARACTER SET utf8mb4;



create table if not exists contacts
(
  id               BIGINT PRIMARY KEY AUTO_INCREMENT,
  open_id          VARCHAR(40)           DEFAULT '' NOT NULL COMMENT '用户OpenId',
  relation_open_id VARCHAR(40)           DEFAULT '' NOT NULL COMMENT '关联的openId',
  `type`           int          not null default 0 comment '类型',
  remark           varchar(255) not null default '' comment '备注',
  created_at       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间戳',
  is_deleted       smallint     not null default 0 comment '是否已删除 0.否 1.是',
  key relationOpenID(open_id,relation_open_id),
  index(is_deleted)
) CHARACTER SET utf8mb4;

