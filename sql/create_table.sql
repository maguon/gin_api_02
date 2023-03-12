SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- CREATE FUNCTION UPDATE_TIMESTAMP_FUNC
create or replace function update_timestamp_func() returns trigger as $$ begin new.updated_at = current_timestamp;
return new;
end $$ language plpgsql;
-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin_info";
CREATE TABLE "public"."admin_info" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL,
    "user_name" character varying(50),
    "password" character varying(200) NOT NULL,
    "avatar" character varying(200) NOT NULL,
    "phone" character varying(50) NOT NULL,
    "email" character varying(100),
    "gender" smallint,
    "type" int4 NOT NULL,
    PRIMARY KEY (id)
);
create trigger admin_info_upt before
update on admin_info for each row execute procedure update_timestamp_func();
select setval('admin_info_id_seq', 1000, false);
-- ----------------------------
-- Table structure for admin_type
-- ----------------------------
CREATE TABLE "public"."admin_type" (
    "id" bigserial NOT NULL,
    "created_on" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_on" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL DEFAULT 1,
    "type_name" character varying(50) NOT NULL,
    "menu_list" jsonb NOT NULL DEFAULT '{}',
    "remark" character varying(400),
    PRIMARY KEY (id)
);
create trigger admin_type_upt before
update on admin_type for each row execute procedure update_timestamp_func();
select setval('admin_type_id_seq', 1000, false);
-- ----------------------------
-- Table structure for app_version
-- ----------------------------
CREATE TABLE public.app_version (
    id bigserial NOT NULL,
    created_on timestamptz NOT NULL DEFAULT now(),
    updated_on timestamptz NOT NULL DEFAULT now(),
    app_type int2 NOT NULL,
    device_type int2 NOT NULL,
    version_ser varchar(20) NULL,
    version_num int4 NULL,
    min_version_num int4 NULL,
    force_up_flag int2 NULL,
    url varchar(200) NULL,
    remarks varchar(200) NULL,
    status int2 NOT NULL,
    CONSTRAINT app_version_pk PRIMARY KEY (id)
);
-- Column comments
COMMENT ON COLUMN public.app_version.app_type IS 'app类型(1-新能源)';
COMMENT ON COLUMN public.app_version.device_type IS '设备类型(1-安卓 2-苹果)';
COMMENT ON COLUMN public.app_version.version_ser IS '版本号';
COMMENT ON COLUMN public.app_version.version_num IS '版本序号';
COMMENT ON COLUMN public.app_version.min_version_num IS '最小支持版本';
COMMENT ON COLUMN public.app_version.force_up_flag IS '是否强制更新(0-不更新 1-更新)';
COMMENT ON COLUMN public.app_version.url IS '下载地址';
COMMENT ON COLUMN public.app_version.remarks IS '备注';
COMMENT ON COLUMN public.app_version.status IS '状态(0:停用,1:启用)';

create trigger app_version_upt before
update on app_version for each row execute procedure update_timestamp_func();
select setval('app_version_id_seq', 1000, false);


-- ----------------------------
-- Table structure for film_info
-- ----------------------------
CREATE TABLE public.film_info (
    id bigserial NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    unique_id varchar(20) NOT NULL,
    url varchar(200) NOT NULL,
    duration int2 NULL,
    thumbnail varchar(200) NULL,
    img varchar(200) NULL,
    studio varchar(200) NULL,
    publish_date date NULL,
    title varchar(50) NULL,
    content varchar(50) NULL,
    category text[] NULL,
    casts text[] NULL,
    remarks varchar(200) NULL,
    status int2 NOT NULL,
    CONSTRAINT film_info_pk PRIMARY KEY (id)
);

create trigger film_info_upt before
update on film_info for each row execute procedure update_timestamp_func();
select setval('film_info_id_seq', 10000, false);
