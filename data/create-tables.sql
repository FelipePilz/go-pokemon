DROP TABLE IF EXISTS `TYPE`;
DROP TABLE IF EXISTS `POKEMON`;
DROP TABLE IF EXISTS `POKEMON_HAS_TYPE`;

USE pokemons;

CREATE TABLE `TYPE`
(
    id   INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128)       NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `POKEMON`
(
    id   INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255)       NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `POKEMON_HAS_TYPE`
(
    pokemon_id INT NOT NULL,
    type_id    INT NOT NULL,
    PRIMARY KEY (`pokemon_id`, `type_id`),
    CONSTRAINT `Pokemon_id_fk` FOREIGN KEY `Pokemon_fk` (`pokemon_id`) REFERENCES `POKEMON` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `Type_id_fk` FOREIGN KEY `Type_fk` (`type_id`) REFERENCES `TYPE` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO `TYPE` (name) VALUES ('Grass'), ('Poison'), ('Fire'), ('Flying'), ('Water');

INSERT INTO `POKEMON` (name) VALUES
('Bulbasaur'), ('Ivysaur'), ('Venusaur'),
('Charmander'), ('Charmeleon'), ('Charizard'),
('Squirtle'), ('Wartortle'), ('Blastoise');

INSERT INTO `POKEMON_HAS_TYPE` (pokemon_id, type_id) VALUES
(1,1), (1,2), (2,1), (2,2), (3,1), (3,2),
(4,3), (5,3), (6,3), (6,4),
(7,5), (8,5), (9,5);
