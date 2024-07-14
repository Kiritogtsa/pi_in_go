-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 03/07/2024 às 02:22
-- Versão do servidor: 10.4.32-MariaDB
-- Versão do PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

-- Configurações de caracteres antigos
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

CREATE DATABASE IF NOT EXISTS `sistrh`;
USE `sistrh`;

-- Estrutura para tabela `trabalhos`
CREATE TABLE `trabalhos` (
  `ID` int(11) NOT NULL,
  `DESCRICAO` varchar(30) DEFAULT NULL,
  `NOME` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Estrutura para tabela `users`
CREATE TABLE `users` (
  `ID` int(11) NOT NULL,
  `NOME` varchar(50) DEFAULT NULL,
  `CPF` varchar(11) DEFAULT NULL UNIQUE,
  `EMAIL` varchar(50) DEFAULT NULL UNIQUE,
  `DATA_NASCIMENTO` date DEFAULT NULL,
  `TELEFONE` varchar(15) DEFAULT NULL UNIQUE,
  `DATA_ADMISSAO` date DEFAULT NULL,
  `SEXO` varchar(10) DEFAULT NULL,
  `SENHA` varchar(60) DEFAULT NULL,
  `GRUPO` varchar(12) NOT NULL DEFAULT 'user',
  `DELETE_AT` DATETIME DEFAULT NULL,
  `TR_ID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Índices de tabela `trabalhos`
ALTER TABLE `trabalhos`
  ADD PRIMARY KEY (`ID`);

-- Índices de tabela `users`
ALTER TABLE `users`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `TR_ID` (`TR_ID`);

-- AUTO_INCREMENT para tabelas `trabalhos`
ALTER TABLE `trabalhos`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT;

-- AUTO_INCREMENT para tabelas `users`
ALTER TABLE `users`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT;

-- Restrições para tabelas `users`
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`TR_ID`) REFERENCES `trabalhos` (`ID`);

COMMIT;

-- Restauração das configurações de caracteres antigos
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
