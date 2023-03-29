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
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
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
-- Table structure for user_info
-- ----------------------------
CREATE TABLE public.user_info (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" int2 NOT NULL,
    "phone" varchar(20) NOT NULL,
    "email" varchar(50) NULL,
    "password" varchar(50) NULL,
    "name" varchar(20) NULL,
    "avatar" varchar(200) NULL,
    "gender" int2 NULL,
    "birth" date NULL,
    "remarks" varchar(200) NULL,
    CONSTRAINT user_info_pk PRIMARY KEY (id)
);
-- Column comments
COMMENT ON COLUMN public.user_info.gender IS '性别(1-男,2-女)';

create trigger user_info_upt before
update on user_info for each row execute procedure update_timestamp_func();
select setval('user_info_id_seq', 100000, false);

-- -------------------------------
-- Table structure for user_device
-- -------------------------------
CREATE TABLE public.user_device (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "last_login_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "user_id" int8 NOT NULL,
    "device_id" varchar(60) NULL,
    "device_token" varchar(60) NULL,
    "device_type" int2 NOT NULL,
    "app_type" int2 NOT NULL,
    "user_type" int2 NOT NULL,
    "version" varchar(20) NULL,
    "version_num" int4 NOT NULL,
    CONSTRAINT user_device_pk PRIMARY KEY (id)
);
-- Column comments
COMMENT ON COLUMN public.user_device.device_id IS '设备标识';
COMMENT ON COLUMN public.user_device.device_token IS '设备推送token';
COMMENT ON COLUMN public.user_device.device_type IS '设备类型(1-安卓 2-苹果)';
COMMENT ON COLUMN public.user_device.user_type IS '设备类型(1-求职 2-招聘)';

create trigger user_device_upt before
update on user_device for each row execute procedure update_timestamp_func();
select setval('user_device_id_seq', 1000, false);


-- ----------------------------
-- Table structure for app_info
-- ----------------------------
CREATE TABLE public.app_info (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" int2 NOT NULL,
    "app_type" int2 NOT NULL,
    "device_type" int2 NOT NULL,
    "version" varchar(20) NULL,
    "version_num" int4 NULL,
    "min_version_num" int4 NULL,
    "force_up_flag" int2 NULL,
    "url" varchar(200) NULL,
    "remarks" varchar(200) NULL,
    CONSTRAINT app_info_pk PRIMARY KEY (id)
);
-- Column comments
COMMENT ON COLUMN public.app_info.app_type IS 'app类型(1-主App)';
COMMENT ON COLUMN public.app_info.device_type IS '设备类型(1-安卓 2-苹果)';
COMMENT ON COLUMN public.app_info.version IS '版本号';
COMMENT ON COLUMN public.app_info.version_num IS '版本序号';
COMMENT ON COLUMN public.app_info.min_version_num IS '最小支持版本';
COMMENT ON COLUMN public.app_info.force_up_flag IS '是否强制更新( 1-更新 0-不更新)';
COMMENT ON COLUMN public.app_info.url IS '下载地址';
COMMENT ON COLUMN public.app_info.remarks IS '备注';
COMMENT ON COLUMN public.app_info.status IS '状态(0:停用,1:启用)';

create trigger app_info_upt before
update on app_info for each row execute procedure update_timestamp_func();
select setval('app_info_id_seq', 1000, false);


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
