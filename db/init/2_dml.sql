use clean_architecture;

SET NAMES utf8mb4;

INSERT INTO `user` (`id`,`name`) VALUES ("1","name1");
INSERT INTO `user` (`id`,`name`) VALUES ("2","name2");
INSERT INTO `user` (`id`,`name`) VALUES ("3","name3");
INSERT INTO `user` (`id`,`name`) VALUES ("4","name4");
INSERT INTO `user` (`id`,`name`) VALUES ("5","name5");
INSERT INTO `user` (`id`,`name`) VALUES ("6","name6");

INSERT INTO `todo` (`todo_id`,`user_id`, `title`, `content`, `status`) VALUES ("1","1","testタイトル","test内容","wip");
INSERT INTO `todo` (`todo_id`,`user_id`, `title`, `content`, `status`) VALUES ("2","1","testタイトル2","test内容2","wip");
INSERT INTO `todo` (`todo_id`,`user_id`, `title`, `content`, `status`) VALUES ("3","2","testタイトル3","test内容3","wip");