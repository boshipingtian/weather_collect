CREATE SCHEMA IF NOT EXISTS weather_colly;

DROP TABLE IF EXISTS country;
CREATE TABLE country
(
    `id`                INT          NOT NULL COMMENT '主键',
    `country_name`      VARCHAR(255) NOT NULL COMMENT '国家名',
    `country_cn_name`   VARCHAR(255) NOT NULL COMMENT '国家中文名',
    `country_full_name` VARCHAR(255) COMMENT '国家全名',
    `created_time`      DATETIME COMMENT '创建时间',
    `updated_time`      DATETIME COMMENT '更新时间',
    PRIMARY KEY (id)
) COMMENT = '国家台账表';
CREATE INDEX unique_country_name ON country (country_name);


DROP TABLE IF EXISTS city;
CREATE TABLE city
(
    `id`           INT          NOT NULL COMMENT '主键',
    `city_name`    VARCHAR(255) NOT NULL COMMENT '城市名',
    `city_type`    INT          NOT NULL COMMENT '城市类型',
    `parent_id`    INT COMMENT '父节点ID',
    `country_id`   INT COMMENT '国家ID',
    `sorts`        INT          NOT NULL COMMENT '排序',
    `created_time` DATETIME COMMENT '创建时间',
    `updated_time` DATETIME COMMENT '更新时间',
    PRIMARY KEY (id)
) COMMENT = '城市台账表';

DROP TABLE IF EXISTS city_type;
CREATE TABLE city_type
(
    `id`           INT          NOT NULL COMMENT '主键',
    `name`         VARCHAR(255) NOT NULL COMMENT '类型',
    `country_id`   INT COMMENT '国家ID',
    `created_time` DATETIME COMMENT '创建时间',
    `updated_time` DATETIME COMMENT '更新时间',
    PRIMARY KEY (Id)
) COMMENT = '城市类型表';

DROP TABLE IF EXISTS weather_city_code;
CREATE TABLE weather_city_code
(
    `id`           INT AUTO_INCREMENT NOT NULL COMMENT '主键',
    `city_id`      VARCHAR(255)       NOT NULL COMMENT '城市名',
    `code`         INT                NOT NULL COMMENT '气象代码',
    `city_pinyin`  VARCHAR(255)       NOT NULL COMMENT '气象城市拼音',
    `created_time` DATETIME COMMENT '创建时间',
    `updated_time` DATETIME COMMENT '更新时间',
    PRIMARY KEY (id)
) COMMENT = '城市气象代码表';
CREATE UNIQUE INDEX UNIQUE_WEATHER_CODE ON weather_city_code (code);
CREATE UNIQUE INDEX UNIQUE_CITY_ID ON weather_city_code (city_id);

DROP TABLE IF EXISTS weather_type;
CREATE TABLE weather_type
(
    `id`           INT          NOT NULL COMMENT '主键',
    `name`         VARCHAR(255) NOT NULL COMMENT '类型',
    `unit`         VARCHAR(255) NOT NULL COMMENT '单位',
    `created_time` DATETIME COMMENT '创建时间',
    `updated_time` DATETIME COMMENT '更新时间',
    PRIMARY KEY (id)
) COMMENT = '气象类型表';

DROP TABLE IF EXISTS weather;
CREATE TABLE weather
(
    `id`           INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `date`         DATE         DEFAULT NULL COMMENT '日期',
    `time`         TIME         DEFAULT NULL COMMENT '时间',
    `city_code`    INT          DEFAULT NULL COMMENT '城市气象代码',
    `type`         INT          DEFAULT NULL COMMENT '气象类型',
    `value`        VARCHAR(255) DEFAULT NULL COMMENT '数据',
    `created_time` DATETIME COMMENT '创建时间',
    `updated_time` DATETIME COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY `unique_city_code_type_date_time` (`city_code`, `type`, `date`, `time`)
) COMMENT = '气象数据表';


