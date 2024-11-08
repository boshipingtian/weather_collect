CREATE SCHEMA WEATHER_COLLY;

DROP TABLE IF EXISTS COUNTRY;
CREATE TABLE COUNTRY
(
    `Id`                INT          NOT NULL COMMENT '主键',
    `COUNTRY_NAME`      VARCHAR(255) NOT NULL COMMENT '国家名',
    `COUNTRY_CN_NAME`   VARCHAR(255) NOT NULL COMMENT '国家中文名',
    `COUNTRY_FULL_NAME` VARCHAR(255) COMMENT '国家全名',
    `CREATED_TIME`      DATETIME COMMENT '创建时间',
    `UPDATED_TIME`      DATETIME COMMENT '更新时间',
    PRIMARY KEY (Id)
) COMMENT = '国家台账表';
CREATE INDEX UNIQUE_COUNTRY_NAME ON COUNTRY (COUNTRY_NAME);


DROP TABLE IF EXISTS CITY;
CREATE TABLE CITY
(
    `Id`           INT          NOT NULL COMMENT '主键',
    `CITY_NAME`    VARCHAR(255) NOT NULL COMMENT '城市名',
    `CITY_TYPE`    INT          NOT NULL COMMENT '城市类型',
    `PARENT_ID`    INT COMMENT '父节点ID',
    `COUNTRY_ID`   INT COMMENT '国家ID',
    `SORTS`        INT          NOT NULL COMMENT '排序',
    `CREATED_TIME` DATETIME COMMENT '创建时间',
    `UPDATED_TIME` DATETIME COMMENT '更新时间',
    PRIMARY KEY (Id)
) COMMENT = '城市台账表';

DROP TABLE IF EXISTS CITY_TYPE;
CREATE TABLE CITY_TYPE
(
    `Id`           INT          NOT NULL COMMENT '主键',
    `NAME`         VARCHAR(255) NOT NULL COMMENT '类型',
    `COUNTRY_ID`   INT COMMENT '国家ID',
    `CREATED_TIME` DATETIME COMMENT '创建时间',
    `UPDATED_TIME` DATETIME COMMENT '更新时间',
    PRIMARY KEY (Id)
) COMMENT = '城市类型表';

DROP TABLE IF EXISTS WEATHER_CITY_CODE;
CREATE TABLE WEATHER_CITY_CODE
(
    `Id`           INT AUTO_INCREMENT NOT NULL COMMENT '主键',
    `CITY_ID`      VARCHAR(255)       NOT NULL COMMENT '城市名',
    `CODE`         INT                NOT NULL COMMENT '气象代码',
    `CITY_PINYIN`  VARCHAR(255)       NOT NULL COMMENT '气象城市拼音',
    `CREATED_TIME` DATETIME COMMENT '创建时间',
    `UPDATED_TIME` DATETIME COMMENT '更新时间',
    PRIMARY KEY (Id)
) COMMENT = '城市气象代码表';
CREATE UNIQUE INDEX UNIQUE_WEATHER_CODE ON WEATHER_CITY_CODE (CODE);
CREATE UNIQUE INDEX UNIQUE_CITY_ID ON WEATHER_CITY_CODE (CITY_ID);

DROP TABLE IF EXISTS WEATHER_TYPE;
CREATE TABLE WEATHER_TYPE
(
    `ID`           INT          NOT NULL COMMENT '主键',
    `NAME`         VARCHAR(255) NOT NULL COMMENT '类型',
    `UNIT`         VARCHAR(255) NOT NULL COMMENT '单位',
    `CREATED_TIME` DATETIME COMMENT '创建时间',
    `UPDATED_TIME` DATETIME COMMENT '更新时间',
    PRIMARY KEY (ID)
) COMMENT = '气象类型表';

DROP TABLE IF EXISTS WEATHER;
CREATE TABLE WEATHER
(
    `ID`           INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `DATE`         DATE         DEFAULT NULL COMMENT '日期',
    `TIME`         TIME         DEFAULT NULL COMMENT '时间',
    `CITY_CODE`    INT          DEFAULT NULL COMMENT '城市气象代码',
    `TYPE`         INT          DEFAULT NULL COMMENT '气象类型',
    `VALUE`        VARCHAR(255) DEFAULT NULL COMMENT '数据',
    `CREATED_TIME` DATETIME COMMENT '创建时间',
    `UPDATED_TIME` DATETIME COMMENT '更新时间',
    PRIMARY KEY (`ID`),
    UNIQUE KEY `WEATHER_CITY_CODE_TYPE_DATE_TIME_UINDEX` (`CITY_CODE`,`TYPE`,`DATE`,`TIME`)
) COMMENT = '气象数据表';


