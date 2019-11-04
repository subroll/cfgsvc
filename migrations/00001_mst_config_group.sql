-- +goose Up
-- +goose StatementBegin
CREATE TABLE `mst_config_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) NOT NULL DEFAULT '',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `mst_config_group`
-- +goose StatementEnd
