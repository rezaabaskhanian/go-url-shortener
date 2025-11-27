
-- +migrate Up
INSERT INTO `permissions` (`id`, `title`) VALUES(1, 'user-list');
INSERT INTO `permissions` (`id`, `title`) VALUES(2, 'user-delete');


-- +migrate Down
DELETE FROM `permissions` WHERE id in (1,2);