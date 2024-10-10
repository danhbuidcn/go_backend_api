-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_crm_user_c (
    `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
    `user_email` varchar(30) NOT NULL COMMENT 'Email',
    `user_phone` varchar(15) NOT NULL COMMENT 'Phone Number',
    `user_username` varchar(30) NOT NULL COMMENT 'Username',
    `user_password` varchar(32) NOT NULL COMMENT 'Password',
    `user_created_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Creation Time',
    `user_updated_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Update Time',
    `user_create_ip_at` varchar(12) NOT NULL COMMENT 'Creation IP',
    `user_last_login_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Last Login Time',
    `user_last_login_ip_at` varchar(12) NOT NULL COMMENT 'Last Login IP',
    `user_login_times` int(11) NOT NULL DEFAULT '0' COMMENT 'Login Times',
    `user_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable',
    PRIMARY KEY (`user_id`),
    KEY `idx_email` (`user_email`),
    KEY `idx_phone` (`user_phone`),
    KEY `idx_username` (`user_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_crm_user_c`;
-- +goose StatementEnd
