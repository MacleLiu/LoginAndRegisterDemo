create table user(
    `user_id` INT UNSIGNED AUTO_INCREMENT,
    `user_name` VARCHAR(10) NOT NULL,
    `password` VARCHAR(40) NOT NULL,
    `salt` VARCHAR(10) NOT NULL
    PRIMARY KEY ( `user_id` ),
    UNIQUE(`user_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;