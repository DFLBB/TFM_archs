-- --------------------------------------------------------
-- Host:                         tfm.cnhygfwtvffy.eu-west-3.rds.amazonaws.com
-- Versión del servidor:         10.4.8-MariaDB-log - Source distribution
-- SO del servidor:              Linux
-- HeidiSQL Versión:             11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Volcando estructura para tabla CARGA_INICIAL_BlockChain.VACUNAS_perros
CREATE TABLE IF NOT EXISTS `VACUNAS_perros` (
  `ID` int(11) NOT NULL,
  `ID_PERRO` int(11) NOT NULL,
  `ID_VETERINARIO` int(11) NOT NULL,
  `COD_VACUNA` varchar(15) COLLATE latin1_spanish_ci NOT NULL,
  `FECHA_ALTA` date NOT NULL,
  `FECHA_BAJA` date DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- Volcando datos para la tabla CARGA_INICIAL_BlockChain.VACUNAS_perros: ~89 rows (aproximadamente)
/*!40000 ALTER TABLE `VACUNAS_perros` DISABLE KEYS */;
INSERT INTO `VACUNAS_perros` (`ID`, `ID_PERRO`, `ID_VETERINARIO`, `COD_VACUNA`, `FECHA_ALTA`, `FECHA_BAJA`) VALUES
	(1, 1, 109, '056969516032348', '2011-09-21', '2012-09-20'),
	(2, 1, 111, '071314377776640', '2012-09-14', '2013-09-14'),
	(3, 1, 112, '443414798419378', '2013-09-25', '2014-09-25'),
	(4, 1, 113, '126274340903642', '2014-09-20', '2015-09-20'),
	(5, 1, 109, '874324442701717', '2015-09-02', '2016-09-01'),
	(6, 1, 117, '750486841192923', '2016-09-03', '2017-09-03'),
	(7, 1, 109, '756117496412357', '2017-09-23', '2018-09-23'),
	(8, 1, 110, '016607690618009', '2018-09-22', '2019-09-22'),
	(9, 1, 116, '358728700651535', '2019-09-03', '2020-09-02'),
	(10, 2, 104, '738527826721341', '2011-09-21', '2012-09-20'),
	(11, 2, 105, '319566998928383', '2012-09-07', '2013-09-07'),
	(12, 2, 113, '908609508451643', '2013-09-03', '2014-09-03'),
	(13, 2, 118, '606601141026804', '2014-09-06', '2015-09-06'),
	(14, 2, 100, '757157947641945', '2015-09-10', '2016-09-09'),
	(15, 2, 116, '542623373253852', '2016-09-13', '2017-09-13'),
	(16, 2, 115, '071196268378940', '2017-09-03', '2018-09-03'),
	(17, 2, 101, '818436938168889', '2018-09-01', '2019-09-01'),
	(18, 2, 114, '554698063577245', '2019-09-07', '2020-09-06'),
	(19, 3, 102, '171309462908013', '2011-09-04', '2012-09-03'),
	(20, 3, 109, '044472933788293', '2012-08-31', '2013-08-31'),
	(21, 3, 103, '161703324096002', '2013-09-24', '2014-09-24'),
	(22, 3, 100, '486287735855475', '2014-09-24', '2015-09-24'),
	(23, 3, 110, '245286606480635', '2015-09-21', '2016-09-20'),
	(24, 3, 115, '104886397817067', '2016-09-13', '2017-09-13'),
	(25, 3, 100, '870380547708255', '2017-09-19', '2018-09-19'),
	(26, 3, 117, '224350503854781', '2018-08-29', '2019-08-29'),
	(27, 3, 104, '429815990319305', '2019-09-24', '2020-09-23'),
	(28, 4, 111, '806653345754979', '2011-09-08', '2012-09-07'),
	(29, 4, 109, '706413263181610', '2012-09-24', '2013-09-24'),
	(30, 4, 116, '462117445154225', '2013-09-23', '2014-09-23'),
	(31, 4, 103, '854017832180502', '2014-09-08', '2015-09-08'),
	(32, 4, 109, '436686122265353', '2015-09-17', '2016-09-16'),
	(33, 4, 106, '042743464971654', '2016-09-21', '2017-09-21'),
	(34, 4, 101, '309749499251833', '2017-09-07', '2018-09-07'),
	(35, 4, 111, '862030646635249', '2018-08-28', '2019-08-28'),
	(36, 4, 106, '129116065920439', '2019-09-01', '2020-08-31'),
	(37, 5, 105, '148442155819799', '2011-09-01', '2012-08-31'),
	(38, 5, 108, '842254616173222', '2012-09-09', '2013-09-09'),
	(39, 5, 101, '115523102800849', '2013-09-20', '2014-09-20'),
	(40, 5, 110, '922879815961123', '2014-09-23', '2015-09-23'),
	(41, 5, 108, '709732404639696', '2015-09-11', '2016-09-10'),
	(42, 5, 115, '184113278225170', '2016-09-03', '2017-09-03'),
	(43, 5, 100, '611847743030496', '2017-09-11', '2018-09-11'),
	(44, 5, 119, '010876111696359', '2018-09-05', '2019-09-05'),
	(45, 5, 111, '229108661626567', '2019-09-03', '2020-09-02'),
	(46, 6, 101, '648449541673483', '2011-09-06', '2012-09-05'),
	(47, 6, 113, '994641591789799', '2012-09-01', '2013-09-01'),
	(48, 6, 112, '624467156477708', '2013-09-23', '2014-09-23'),
	(49, 6, 107, '565984231946976', '2014-08-30', '2015-08-30'),
	(50, 6, 109, '132522152860204', '2015-09-21', '2016-09-20'),
	(51, 6, 113, '790385356909024', '2016-09-22', '2017-09-22'),
	(52, 6, 111, '124844765406889', '2017-09-13', '2018-09-13'),
	(53, 6, 102, '005178038035686', '2018-09-11', '2019-09-11'),
	(54, 6, 105, '185603798539941', '2019-09-09', '2020-09-08'),
	(55, 7, 114, '533698783753160', '2011-09-10', '2012-09-09'),
	(56, 7, 102, '742246013825989', '2012-09-05', '2013-09-05'),
	(57, 7, 116, '504557018638176', '2013-09-14', '2014-09-14'),
	(58, 7, 114, '500095286872327', '2014-08-29', '2015-08-29'),
	(59, 7, 105, '995306360530953', '2015-09-11', '2016-09-10'),
	(60, 7, 115, '271852702155573', '2016-09-21', '2017-09-21'),
	(61, 7, 104, '646232975317383', '2017-09-11', '2018-09-11'),
	(62, 7, 116, '812400230031834', '2018-09-09', '2019-09-09'),
	(63, 7, 105, '064091853857126', '2019-09-19', '2020-09-18'),
	(64, 8, 112, '769171176263291', '2011-09-02', '2012-09-01'),
	(65, 8, 114, '092442995954718', '2012-09-12', '2013-09-12'),
	(66, 8, 112, '118032130522684', '2013-09-26', '2014-09-26'),
	(67, 8, 105, '804171886112702', '2014-09-22', '2015-09-22'),
	(68, 8, 105, '255763546801884', '2015-09-09', '2016-09-08'),
	(69, 8, 107, '001532781870601', '2016-09-05', '2017-09-05'),
	(70, 8, 102, '514497165116014', '2017-09-08', '2018-09-08'),
	(71, 8, 109, '379834508877094', '2018-08-28', '2019-08-28'),
	(72, 8, 104, '355316399927546', '2019-09-03', '2020-09-02'),
	(73, 9, 119, '090206662276953', '2011-09-27', '2012-09-26'),
	(74, 9, 107, '888784924418466', '2012-09-08', '2013-09-08'),
	(75, 9, 106, '186456416907215', '2013-09-21', '2014-09-21'),
	(76, 9, 114, '004285865467308', '2014-09-13', '2015-09-13'),
	(77, 9, 114, '066455885829829', '2015-09-04', '2016-09-03'),
	(78, 9, 111, '061593633202458', '2016-09-25', '2017-09-25'),
	(79, 9, 106, '880591801256493', '2017-09-03', '2018-09-03'),
	(80, 9, 106, '460724289026786', '2018-09-18', '2019-09-18'),
	(81, 9, 106, '918594998231712', '2019-09-03', '2020-09-02'),
	(82, 10, 107, '405161239584127', '2005-05-14', '2006-05-14'),
	(83, 10, 107, '958048423713119', '2006-05-08', '2007-05-08'),
	(84, 10, 115, '901318535116798', '2007-05-01', '2008-04-30'),
	(85, 10, 116, '626599984827079', '2008-05-13', '2009-05-13'),
	(86, 10, 112, '199598794057592', '2009-05-18', '2010-05-18'),
	(87, 10, 106, '974183587333359', '2010-05-24', '2011-05-24'),
	(88, 10, 102, '969955123932991', '2011-05-26', '2012-05-25'),
	(89, 10, 106, '307606693643711', '2012-05-01', '2013-05-01');
/*!40000 ALTER TABLE `VACUNAS_perros` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
