create database if not exists im_server_db;
use im_server_db;

create table if not exists chat
(
    id                bigint unsigned auto_increment primary key,
    created_at        datetime(3)     null,
    updated_at        datetime(3)     null,
    chat_sender_id    bigint unsigned null,
    chat_recv_user_id bigint unsigned null,
    chat_msg_type     tinyint         null,
    chat_msg_preview  varchar(64)     null,
    chat_msg          longtext        null,
    chat_sys_msg      longtext        null
);

create table if not exists `group`
(
    id                    bigint unsigned auto_increment primary key,
    created_at            datetime(3)     null,
    updated_at            datetime(3)     null,
    group_name            varchar(32)     null,
    group_abstract        varchar(128)    null,
    group_avatar          varchar(256)    null,
    group_creator         bigint unsigned null,
    group_size            bigint          null,
    group_is_search       tinyint(1)      null,
    group_is_invite       tinyint(1)      null,
    group_is_tmp_session  tinyint(1)      null,
    group_is_prohibit     tinyint(1)      null,
    group_verify          tinyint         null,
    group_verify_question longtext        null
);

create table if not exists chat_member
(
    group_id         bigint unsigned null,
    user_id          bigint unsigned null,
    member_name      varchar(32)     null,
    role             bigint          null,
    prohibition_time bigint          null,
    constraint fk_chat_member_group_model foreign key (group_id) references `group` (id)
);

create table if not exists group_msg
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)     null,
    updated_at   datetime(3)     null,
    group_id     bigint unsigned null,
    send_user_id bigint unsigned null,
    msg_preview  longtext        null,
    msg_type     tinyint         null,
    msg          longtext        null,
    system_msg   longtext        null,
    constraint fk_group_msg_group_model foreign key (group_id) references `group` (id)
);

create table if not exists group_verify
(
    id                    bigint unsigned auto_increment primary key,
    created_at            datetime(3)     null,
    updated_at            datetime(3)     null,
    group_id              bigint unsigned null,
    user_id               bigint unsigned null,
    status                tinyint         null,
    type                  tinyint         null,
    additional_messages   varchar(32)     null,
    verification_question longtext        null,
    constraint fk_group_verify_group_model foreign key (group_id) references `group` (id)
);

create table if not exists user
(
    id            bigint unsigned auto_increment primary key,
    created_at    datetime(3)  null,
    updated_at    datetime(3)  null,
    user_nickname varchar(32)  null,
    user_pwd      varchar(64)  null,
    user_abstract varchar(128) null,
    user_avatar   varchar(256) null,
    user_ip       varchar(32)  null,
    user_addr     varchar(64)  null
);

create table if not exists friend_verify
(
    id                    bigint unsigned auto_increment primary key,
    created_at            datetime(3)     null,
    updated_at            datetime(3)     null,
    send_user_id          bigint unsigned null,
    recv_user_id          bigint unsigned null,
    status                tinyint         null,
    additional_msg        varchar(128)    null,
    verification_question longtext        null,
    constraint fk_friend_verify_recv_user foreign key (recv_user_id) references user (id),
    constraint fk_friend_verify_send_user foreign key (send_user_id) references user (id)
);

create table if not exists friends
(
    id           bigint unsigned auto_increment primary key,
    created_at   datetime(3)     null,
    updated_at   datetime(3)     null,
    send_user_id bigint unsigned null,
    recv_user_id bigint unsigned null,
    notice       varchar(128)    null,
    constraint fk_friends_recv_user foreign key (recv_user_id) references user (id),
    constraint fk_friends_send_user foreign key (send_user_id) references user (id)
);

create table if not exists user_conf
(
    id                      bigint unsigned auto_increment primary key,
    created_at              datetime(3)     null,
    updated_at              datetime(3)     null,
    user_id                 bigint unsigned null,
    recall_msg              varchar(32)     null,
    is_friend_online_notify tinyint(1)      null,
    is_online               tinyint(1)      null,
    is_sound                tinyint(1)      null,
    is_secure_link          tinyint(1)      null,
    is_save_pwd             tinyint(1)      null,
    search_user             tinyint         null,
    verification            tinyint         null,
    verification_question   longtext        null,
    constraint fk_user_conf_user_model foreign key (user_id) references user (id)
);

