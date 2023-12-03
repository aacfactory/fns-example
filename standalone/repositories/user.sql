CREATE TABLE "FNS"."USER"
(
    "ID"        character varying(63)  NOT NULL PRIMARY KEY,
    "CREATE_BY" character varying(63)  NOT NULL DEFAULT '':: character varying,
    "CREATE_AT" timestamp with time zone,
    "MODIFY_BY" character varying(63)  NOT NULL DEFAULT '':: character varying,
    "MODIFY_AT" timestamp with time zone,
    "DELETE_BY" character varying(63)  NOT NULL DEFAULT '':: character varying,
    "DELETE_AT" timestamp with time zone,
    "VERSION"   bigint                 NOT NULL DEFAULT 0,
    "NICKNAME"  character varying(255) NOT NULL,
    "MOBILE"    character varying(255) NOT NULL DEFAULT '':: character varying,
    "GENDER"    character varying(63)  NOT NULL DEFAULT 'N':: character varying,
    "BIRTHDAY"  timestamp with time zone,
    "AVATAR"    jsonb                  NOT NULL DEFAULT '{}'::jsonb
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS "FNS"."USER"
    OWNER to aacfactory;