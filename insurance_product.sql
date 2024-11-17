-- phpMyAdmin SQL Dump
-- version 5.2.1deb3
-- https://www.phpmyadmin.net/
--
-- Хост: localhost:3306
-- Время создания: Ноя 17 2024 г., 08:40
-- Версия сервера: 10.11.8-MariaDB-0ubuntu0.24.04.1
-- Версия PHP: 8.3.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `insurance_product`
--

-- --------------------------------------------------------

--
-- Структура таблицы `all_user_products`
--

CREATE TABLE `all_user_products` (
  `id` int(11) NOT NULL,
  `admin` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `all_user_products_products`
--

CREATE TABLE `all_user_products_products` (
  `all_user_products_id` int(11) DEFAULT NULL,
  `new_product_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `all_user_products_products_patterns`
--

CREATE TABLE `all_user_products_products_patterns` (
  `all_user_products_id` int(11) DEFAULT NULL,
  `new_product_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `new_products`
--

CREATE TABLE `new_products` (
  `id` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `product_param` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `parameters`
--

CREATE TABLE `parameters` (
  `id` int(11) NOT NULL,
  `name` longtext NOT NULL,
  `type` longtext NOT NULL,
  `default_value` longtext DEFAULT NULL,
  `dictionary_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `parameters`
--

INSERT INTO `parameters` (`id`, `name`, `type`, `default_value`, `dictionary_id`) VALUES
(5, 'MAKSIM', 'person', '10', 0);

-- --------------------------------------------------------

--
-- Структура таблицы `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `update_date` date NOT NULL,
  `version_description` text DEFAULT NULL,
  `series_prefix` varchar(255) DEFAULT NULL,
  `series_postfix` varchar(255) DEFAULT NULL,
  `number_prefix` varchar(255) DEFAULT NULL,
  `number_postfix` varchar(255) DEFAULT NULL,
  `numerator` varchar(255) DEFAULT NULL,
  `custom_number_method` varchar(255) DEFAULT NULL,
  `individual_parameters` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`individual_parameters`)),
  `cost_formula` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `product_patterns` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `new_products` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `all_products` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `users_reg`
--

CREATE TABLE `users_reg` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `all_user_products`
--
ALTER TABLE `all_user_products`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `all_user_products_products`
--
ALTER TABLE `all_user_products_products`
  ADD KEY `all_user_products_id` (`all_user_products_id`),
  ADD KEY `new_product_id` (`new_product_id`);

--
-- Индексы таблицы `all_user_products_products_patterns`
--
ALTER TABLE `all_user_products_products_patterns`
  ADD KEY `all_user_products_id` (`all_user_products_id`),
  ADD KEY `new_product_id` (`new_product_id`);

--
-- Индексы таблицы `new_products`
--
ALTER TABLE `new_products`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `parameters`
--
ALTER TABLE `parameters`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Индексы таблицы `users_reg`
--
ALTER TABLE `users_reg`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `all_user_products`
--
ALTER TABLE `all_user_products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `new_products`
--
ALTER TABLE `new_products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `parameters`
--
ALTER TABLE `parameters`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT для таблицы `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `users_reg`
--
ALTER TABLE `users_reg`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `all_user_products_products`
--
ALTER TABLE `all_user_products_products`
  ADD CONSTRAINT `all_user_products_products_ibfk_1` FOREIGN KEY (`all_user_products_id`) REFERENCES `all_user_products` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `all_user_products_products_ibfk_2` FOREIGN KEY (`new_product_id`) REFERENCES `new_products` (`id`) ON DELETE CASCADE;

--
-- Ограничения внешнего ключа таблицы `all_user_products_products_patterns`
--
ALTER TABLE `all_user_products_products_patterns`
  ADD CONSTRAINT `all_user_products_products_patterns_ibfk_1` FOREIGN KEY (`all_user_products_id`) REFERENCES `all_user_products` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `all_user_products_products_patterns_ibfk_2` FOREIGN KEY (`new_product_id`) REFERENCES `new_products` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
