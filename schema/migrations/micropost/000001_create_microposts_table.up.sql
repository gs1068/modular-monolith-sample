CREATE TABLE microposts (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    content TEXT NOT NULL,
    user_id int(10) unsigned NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT 'マイクロポストテーブル';
