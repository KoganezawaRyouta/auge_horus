USE coincheck;

CREATE TABLE `tickers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `last` int(18) NOT NULL,
  `bid` int(18) NOT NULL,
  `ask` int(18) NOT NULL,
  `high` int(18) NOT NULL,
  `low` int(18) NOT NULL,
  `volume` float NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
;

CREATE TABLE `trades` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trade_id` int(18) NOT NULL,
  `amount` float NOT NULL,
  `rate` bigint(20) NOT NULL,
  `order_type` varchar(256) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
;
