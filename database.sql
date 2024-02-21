SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

CREATE TABLE `products` (
  `id` smallint(5) UNSIGNED NOT NULL,
  `category` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `picture` varchar(255) NOT NULL,
  `price` mediumint(9) NOT NULL,
  `stock` mediumint(9) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_hungarian_ci;

--
-- A tábla adatainak kiíratása `products`
--

INSERT INTO `products` (`id`, `category`, `name`, `description`, `picture`, `price`, `stock`) VALUES
(1, 'Bogyós', 'Málna', 'Kézzel termelt egészség', 'https://hur.webmania.cc/img/malna.jpg', 3800, 500),
(2, 'Bogyós', 'Áfonya', 'Az erdő kincse az otthonodba', 'https://hur.webmania.cc/img/afonya.jpg', 3250, 120),
(3, 'Bogyós', 'Szeder', 'A hagyományos csemege', 'https://hur.webmania.cc/img/szeder.jpg', 1700, 40),
(4, 'Bogyós', 'Eper', 'Egy tavaszi harapás', 'https://hur.webmania.cc/img/eper.jpg', 1440, 0),
(5, 'Bogyós', 'Homoktövis', 'Mezei csemege', 'https://hur.webmania.cc/img/homoktovis.jpg', 3200, 100),
(6, 'Bogyós', 'Som', 'A fanyar gyönyör', 'https://hur.webmania.cc/img/som.jpg', 900, 10),
(7, 'Bogyós', 'Fanyarka', 'Édes mint a méz', 'https://hur.webmania.cc/img/fanyarka.jpg', 990, 25),
(8, 'Bogyós', 'Piszke', 'Egres', 'https://hur.webmania.cc/img/piszke.jpg', 750, 100),
(9, 'Bogyós', 'Ribizli', 'Fanyar, vasban gazdag', 'https://hur.webmania.cc/img/ribizli.jpg', 1300, 170),
(10, 'Magyaros', 'Meggy', 'A falusi kincs', 'https://hur.webmania.cc/img/meggy.jpg', 600, 300),
(11, 'Magyaros', 'Cseresznye', 'A falusi kincs', 'https://hur.webmania.cc/img/cseresznye.jpg', 900, 300),
(12, 'Magyaros', 'Szilva', 'A falusi kincs', 'https://hur.webmania.cc/img/szilva.jpg', 770, 200);

CREATE TABLE `users` (
  `id` smallint(5) UNSIGNED NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(32) NOT NULL,
  `token` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- A tábla adatainak kiíratása `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `token`) VALUES
(1, 'rrd@webmania.cc', MD5('rrd'), '270fsdg04%rt2f6$)b4eblok0dfgauranga'),

ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

ALTER TABLE `products`
  MODIFY `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

ALTER TABLE `users`
  MODIFY `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;
