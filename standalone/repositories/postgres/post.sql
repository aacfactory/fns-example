CREATE TABLE "FNS"."POST"
(
    "ID"        character varying(63)  NOT NULL PRIMARY KEY,
    "VERSION"   bigint                 NOT NULL DEFAULT 0,
    "USER_ID"   character varying(63)  NOT NULL,
    "CREATE_AT" timestamp with time zone,
    "TITLE"     character varying(255) NOT NULL DEFAULT '':: character varying,
    "CONTENT"   character varying(512) NOT NULL DEFAULT '':: character varying
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS "FNS"."POST"
    OWNER to aacfactory;


CREATE SEQUENCE IF NOT EXISTS "FNS"."POST_COMMENT_ID"
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE "FNS"."POST_COMMENT_ID"
    OWNER TO aacfactory;

CREATE TABLE "FNS"."POST_COMMENT"
(
    "ID"        bigint                 NOT NULL DEFAULT NEXTVAL('"FNS"."POST_COMMENT_ID"') PRIMARY KEY,
    "POST_ID"   character varying(63)  NOT NULL,
    "USER_ID"   character varying(63)  NOT NULL,
    "CREATE_AT" timestamp with time zone,
    "CONTENT"   character varying(512) NOT NULL DEFAULT '':: character varying
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS "FNS"."POST_COMMENT"
    OWNER to aacfactory;


CREATE SEQUENCE IF NOT EXISTS "FNS"."POST_LIKE_ID"
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE "FNS"."POST_LIKE_ID"
    OWNER TO aacfactory;

CREATE TABLE "FNS"."POST_LIKE"
(
    "ID"      bigint                NOT NULL DEFAULT NEXTVAL('"FNS"."POST_LIKE_ID"') PRIMARY KEY,
    "POST_ID" character varying(63) NOT NULL,
    "USER_ID" character varying(63) NOT NULL
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS "FNS"."POST_LIKE"
    OWNER to aacfactory;