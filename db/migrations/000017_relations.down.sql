ALTER TABLE `consultants`
DROP FOREIGN KEY `fk_consultants_user`;

ALTER TABLE `consultants`
DROP COLUMN `user_id`;
