SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `clean_architecture` DEFAULT CHARACTER SET utf8mb4 ;
USE `clean_architecture` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `clean_architecture`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clean_architecture`.`user` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- Table `clean_architecture`.`todo`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clean_architecture`.`todo` (
  `todo_id` VARCHAR(128) NOT NULL COMMENT 'todo ID',
  `user_id` VARCHAR(128) NOT NULL COMMENT 'ユーザ ID',
  `title` VARCHAR(128) NOT NULL COMMENT 'タイトル',
  `content` VARCHAR(128) NOT NULL COMMENT '内容',
  `status` VARCHAR(128) NOT NULL COMMENT 'ステータス',
  PRIMARY KEY (`todo_id`),
  FOREIGN KEY (`user_id`)
    REFERENCES user(`id`)
    ON DELETE CASCADE
  )
ENGINE = InnoDB
COMMENT = 'todo';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
