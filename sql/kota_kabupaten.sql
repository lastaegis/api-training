create table KOTA_KABUPATEN
(
    ID             int auto_increment
        primary key,
    ID_PROVINSI    int unsigned null,
    KOTA_KABUPATEN varchar(255) null,
    CREATED_AT     timestamp    null,
    CREATED_BY     varchar(255) null,
    UPDATED_AT     timestamp    null,
    UPDATED_BY     varchar(255) null,
    DELETED_AT     timestamp    null,
    DELETED_BY     varchar(255) null,
    SYNC_STATUS    tinyint      null,
    constraint FK_KOTA_KABUPATEN
        foreign key (ID_PROVINSI) references PROVINSI (ID)
);