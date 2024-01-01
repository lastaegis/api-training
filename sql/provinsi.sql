create table if not exists PROVINSI
(
    ID          int auto_increment
        primary key,
    PROVINSI    varchar(255) null,
    CREATED_AT  datetime     null,
    CREATED_BY  varchar(255) null,
    UPDATED_AT  datetime     null,
    UPDATED_BY  varchar(255) null,
    DELETE_AT   datetime     null,
    DELETED_BY  varchar(255) null,
    SYNC_STATUS tinyint      null
);