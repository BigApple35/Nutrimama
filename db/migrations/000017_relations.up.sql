ALTER TABLE `consultants`
ADD COLUMN `user_id` bigint UNSIGNED NOT NULL;

ALTER TABLE `consultants`
ADD CONSTRAINT `fk_consultants_user`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`)
    ON DELETE CASCADE ON UPDATE CASCADE;
