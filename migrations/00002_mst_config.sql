-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_config` (
  `id` int(19) unsigned NOT NULL AUTO_INCREMENT,
  `mst_config_group_id` int(10) unsigned NOT NULL,
  `key` varchar(255) NOT NULL DEFAULT '',
  `value` varchar(1024) NOT NULL DEFAULT '',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `mst_config_group_id` (`mst_config_group_id`),
  CONSTRAINT `mst_config_ibfk_1` FOREIGN KEY (`mst_config_group_id`) REFERENCES `mst_config_group` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=latin1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_config`
-- +goose StatementEnd
